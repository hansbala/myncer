import { z } from 'zod'
import { createTRPCRouter, protectedProcedure } from '../trpc'
import { DATASOURCE } from '~/core/datasources'
import { SYNC_FREQUENCY } from '~/core/syncFrequency'

export const syncsRouter = createTRPCRouter({
  getSyncs: protectedProcedure.query(async ({ ctx }) => {
    return await ctx.db.sync.findMany({
      where: {
        userId: ctx.session.user.id,
      },
    })
  }),
  createNewSync: protectedProcedure
    .input(
      z.object({
        sourcePlaylist: z.object({
          datasource: z.nativeEnum(DATASOURCE),
          playlistId: z.string(),
        }),
        destinationPlaylist: z.object({
          datasource: z.nativeEnum(DATASOURCE),
          playlistId: z.string(),
        }),
        syncFrequency: z.nativeEnum(SYNC_FREQUENCY),
        useWebhooks: z.boolean(),
      })
    )
    .mutation(async ({ ctx, input }) => {
      await ctx.db.sync.create({
        data: {
          sourcePlaylistDatasource: input.sourcePlaylist.datasource,
          destinationPlaylistDatasource: input.destinationPlaylist.datasource,
          sourcePlaylistId: input.sourcePlaylist.playlistId,
          destinationPlaylistId: input.destinationPlaylist.playlistId,
          syncFrequency: input.syncFrequency,
          useWebhooks: input.useWebhooks,
          userId: ctx.session.user.id,
        },
      })
    }),
})
