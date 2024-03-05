import { authUserForGoogleFirstTime, getGoogleAuthorizationUrl } from "~/server/utils/google";
import { createTRPCRouter, protectedProcedure } from "../trpc";
import { z } from "zod";

export const googleRouter = createTRPCRouter({
  getAuthorizationUrl: protectedProcedure.query(() => {
    return getGoogleAuthorizationUrl()
  }),
  setAuthorizationCode: protectedProcedure.input(z.object({ authorizationCode: z.string() })).mutation(async ({ ctx, input }) => {
    const { accessToken, refreshToken } = await authUserForGoogleFirstTime(input.authorizationCode)
    await ctx.db.googleKey.create({
      data: {
        authCode: input.authorizationCode,
        accessCode: accessToken,
        refreshToken,
        userId: ctx.session.user.id
      }
    })
  }),
  isAuthenticated: protectedProcedure.query(async ({ ctx }) => {
    const foundKey = await ctx.db.googleKey.findUnique({
      where: {
        userId: ctx.session.user.id
      }
    })
    return foundKey ? true : false
  })
})