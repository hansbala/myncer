import { createTRPCRouter } from "~/server/api/trpc";
import { spotifyRouter } from "./routers/spotifyRouter";
import { googleRouter } from "./routers/googleRouter";
import { syncsRouter } from "./routers/syncsRouter";

/**
 * This is the primary router for your server.
 *
 * All routers added in /api/routers should be manually added here.
 */
export const appRouter = createTRPCRouter({
  spotify: spotifyRouter,
  google: googleRouter,
  syncs: syncsRouter
});

// export type definition of API
export type AppRouter = typeof appRouter;
