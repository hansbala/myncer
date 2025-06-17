import type { OneWaySync } from "@/generated_api/src"
import { usePlaylist } from "@/hooks/usePlaylist"
import { YoutubeIcon, Music } from "lucide-react"
import { cn } from "@/lib/utils"

export const OneWaySyncRender = ({ sync }: { sync: OneWaySync }) => {
  const { playlist: sourcePlaylist, loading: loadingSource } = usePlaylist({
    datasource: sync.source.datasource,
    playlistId: sync.source.playlistId,
  })

  const { playlist: destinationPlaylist, loading: loadingDest } = usePlaylist({
    datasource: sync.destination.datasource,
    playlistId: sync.destination.playlistId,
  })

  const renderDatasource = (
    datasource: string,
    playlistName?: string,
    loading?: boolean
  ) => {
    const name = datasource.toLowerCase()
    const isSpotify = name === "spotify"
    const isYoutube = name === "youtube"

    const Icon = isSpotify ? Music : isYoutube ? YoutubeIcon : undefined
    const brandColor = isSpotify
      ? "text-green-500"
      : isYoutube
        ? "text-red-500"
        : "text-gray-500"

    return (
      <div className="flex flex-col w-full max-w-[240px] space-y-2">
        <div className="flex items-center gap-2">
          {Icon && <Icon className={cn("w-4 h-4", brandColor)} />}
          <span className="capitalize font-medium">{datasource}</span>
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
        {renderDatasource(sync.source.datasource, sourcePlaylist?.name, loadingSource)}

        <div className="text-gray-400 text-lg">→</div>

        <div className="justify-self-end">
          {renderDatasource(sync.destination.datasource, destinationPlaylist?.name, loadingDest)}
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

