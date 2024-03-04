import { authUserForSpotifyFirstTime, createSpotifyAuthUrl, getCurrentUserPlaylists } from "~/server/utils/spotify"
import { type createTRPCContext, createTRPCRouter, protectedProcedure } from "../trpc"
import { z } from "zod"

const _getAuthorizationCode = async (ctx: Awaited<ReturnType<typeof createTRPCContext>>): Promise<string> => {
  if (!ctx.session) {
    throw new Error('No session found')
  }
  const accessToken = await ctx.db.spotifyApiKey.findUnique({
    where: {
      userId: ctx.session.user.id
    },
    select: {
      authCode: true
    }
  })
  if (!accessToken) {
    throw new Error('Failed to get access token of spotify for user')
  }
  return accessToken.authCode
}

export const spotifyRouter = createTRPCRouter({
  getAuthorizationUrl: protectedProcedure.query(async ({ ctx }) => {
    return createSpotifyAuthUrl()
  }),
  setAuthorizationCode: protectedProcedure.input(z.object({ authCode: z.string() })).mutation(async ({ ctx, input }) => {
    if (!ctx.session) {
      throw new Error('No user session found')
    }
    const { accessToken, refreshToken } = await authUserForSpotifyFirstTime(input.authCode)
    await ctx.db.spotifyApiKey.create({
      data: {
        authCode: input.authCode,
        accessToken,
        refreshToken,
        userId: ctx.session.user.id
      }
    })
  }),
  getAuthorizationCode: protectedProcedure.query(async ({ ctx }) => {
    return _getAuthorizationCode(ctx)
  }),
  getCurrentUserPlaylists: protectedProcedure.query(async ({ ctx }) => {
    if (!ctx.session) {
      throw new Error('No user session found')
    }
    return await getCurrentUserPlaylists(ctx.session.user.id)
  }),
})