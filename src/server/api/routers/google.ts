import { getGoogleAuthorizationUrl } from "~/server/utils/google";
import { createTRPCRouter, protectedProcedure } from "../trpc";

export const googleRouter = createTRPCRouter({
  getAuthorizationUrl: protectedProcedure.query(() => {
    return getGoogleAuthorizationUrl()
  })
})