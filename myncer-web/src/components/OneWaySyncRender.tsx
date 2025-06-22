import { usePlaylist } from "@/hooks/usePlaylist"
import { YoutubeIcon, Music, ArrowRight } from "lucide-react"
import { cn, getDatasourceLabel } from "@/lib/utils"
import type { OneWaySync } from "@/generated_grpc/myncer/sync_pb"
import { Datasource } from "@/generated_grpc/myncer/datasource_pb"

export const OneWaySyncRender = ({ sync }: { sync: OneWaySync }) => {
  const { playlist: sourcePlaylist, loading: loadingSource } = usePlaylist({
    datasource: sync.source?.datasource,
    playlistId: sync.source?.playlistId,
  })

  const { playlist: destinationPlaylist, loading: loadingDest } = usePlaylist({
    datasource: sync.destination?.datasource,
    playlistId: sync.destination?.playlistId,
  })

  const getDatasourceIcon = (datasource: Datasource) => {
    switch (datasource) {
      case Datasource.SPOTIFY:
        return Music
      case Datasource.YOUTUBE:
        return YoutubeIcon
      default:
        return Music
    }
  }

  const getDatasourceBrandColor = (datasource: Datasource) => {
    switch (datasource) {
      case Datasource.SPOTIFY:
        return "text-green-500"
      case Datasource.YOUTUBE:
        return "text-red-500"
      default:
        return "text-gray-500"
    }
  }

  const renderDatasource = (
    datasource: Datasource,
    playlistName?: string,
    loading?: boolean
  ) => {
    const Icon = getDatasourceIcon(datasource)
    const brandColor = getDatasourceBrandColor(datasource)
    return (
      <div className="flex flex-col w-full max-w-[240px] space-y-2">
        <div className="flex items-center gap-2">
          {Icon && <Icon className={cn("w-4 h-4", brandColor)} />}
          <span className="capitalize font-medium">{getDatasourceLabel(datasource)}</span>
        </div>
        <div className="text-xs text-muted-foreground truncate w-full">
          {loading ? "Loading..." : playlistName ?? "Unnamed playlist"}
        </div>
      </div>
    )
  }

  return (
    <div className="space-y-2 text-sm w-full">
      <div className="grid grid-cols-[1fr_auto_1fr] items-center gap-4">
        {renderDatasource(
          sync.source?.datasource ?? Datasource.UNSPECIFIED,
          sourcePlaylist?.name ?? "Unnamed playlist",
          loadingSource,
        )}
        <ArrowRight />
        <div className="justify-self-end">
          {renderDatasource(
            sync.destination?.datasource ?? Datasource.UNSPECIFIED,
            destinationPlaylist?.name ?? "Unnamed playlist",
            loadingDest,
          )}
        </div>
      </div>

      {sync.overwriteExisting && (
        <div className="text-xs text-yellow-800 bg-yellow-100 inline-block px-2 py-0.5 rounded">
          Overwrites destination playlist
        </div>
      )}
    </div>
  )
}

