import { createTRPCRouter, protectedProcedure } from "../trpc";

export const googleRouter = createTRPCRouter({
  getAuthorizationUrl: protectedProcedure.query(() => {
    return 'hello'
  })
})