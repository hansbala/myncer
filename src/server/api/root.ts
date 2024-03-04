import { createTRPCRouter } from "~/server/api/trpc";
import { secretsRouter } from "./routers/secrets";
import { spotifyRouter } from "./routers/spotify";

/**
 * This is the primary router for your server.
 *
 * All routers added in /api/routers should be manually added here.
 */
export const appRouter = createTRPCRouter({
  secrets: secretsRouter,
  spotify: spotifyRouter
});

// export type definition of API
export type AppRouter = typeof appRouter;
