import { decryptUsingEnvKey, encryptUsingEnvKey } from "~/server/utils/crypto";
import { createTRPCRouter, protectedProcedure } from "../trpc";

import { z } from "zod";

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
        const userId = ctx.session.user.id
        const spotifyApiKey = await ctx.db.spotifyApiKey.findUnique({
            where: {
                userId
            }
        })
        if (!spotifyApiKey) {
            return null
        }
        const { clientId, iv, encryptedClientSecret } = spotifyApiKey
        const clientSecret = decryptUsingEnvKey(encryptedClientSecret, iv)
        return {
            clientId,
            clientSecret
        }
    })
})