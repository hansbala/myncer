import { z } from "zod";
import { createTRPCRouter, protectedProcedure } from "../trpc";
import { DATASOURCE } from "~/common/datasources";
import { SYNC_FREQUENCY } from "~/common/syncFrequency";

export const syncsRouter = createTRPCRouter({
  createNewSync: protectedProcedure.input(z.object({
    sourcePlaylist: z.object({
      datasource: z.nativeEnum(DATASOURCE),
      playlistId: z.string()
    }),
    destinationPlaylist: z.object({
      datasource: z.nativeEnum(DATASOURCE),
      playlistId: z.string()
    }),
    syncFrequency: z.nativeEnum(SYNC_FREQUENCY),
    useWebhooks: z.boolean()
  })).mutation(async ({ ctx, input }) => {
    throw new Error('Not implemented')
  })
})