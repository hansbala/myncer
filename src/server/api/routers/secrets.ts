import { decryptUsingEnvKey, encryptUsingEnvKey } from "~/server/utils/crypto";
import { type createTRPCContext, createTRPCRouter, protectedProcedure } from "../trpc";

import { z } from "zod";
import { getSpotifyAccessToken } from "~/server/utils/spotify";

const _getSpotifySecret = async (ctx: Awaited<ReturnType<typeof createTRPCContext>>) => {
    if (!ctx.session) {
        throw new Error('No session found')
    }
    const userId = ctx.session.user.id
    const spotifyApiKey = await ctx.db.spotifyApiKey.findUnique({
        where: {
            userId
        }
    })
    if (!spotifyApiKey) {
        throw new Error('No Spotify API key found')
    }
    const { clientId, iv, encryptedClientSecret } = spotifyApiKey
    const clientSecret = decryptUsingEnvKey(encryptedClientSecret, iv)
    return {
        clientId,
        clientSecret
    }
}

export const secretsRouter = createTRPCRouter({
    setSpotifySecret: protectedProcedure.input(z.object({ clientId: z.string(), clientSecret: z.string() }))
        .mutation(async ({ ctx, input }) => {
            const { clientId, clientSecret } = input
            const { iv, encryptedText: encryptedClientSecret } = encryptUsingEnvKey(clientSecret)

            const userId = ctx.session.user.id
            await ctx.db.spotifyApiKey.upsert({
                where: {
                    userId
                },
                create: {
                    clientId,
                    iv,
                    encryptedClientSecret,
                    userId
                },
                update: {
                    clientId,
                    iv,
                    encryptedClientSecret,
                }
            })
        }),
    getSpotifySecret: protectedProcedure.query(async ({ ctx }) => {
        return _getSpotifySecret(ctx)
    }),
    fetchSpotifyAccessToken: protectedProcedure.mutation(async ({ ctx }) => {
        // TODO: Cache and use exisiting access token if it exists + timestamp is valid
        // for now, always fetch a new one
        const { clientId, clientSecret } = await _getSpotifySecret(ctx)
        return await getSpotifyAccessToken(clientId, clientSecret)
    }),
})