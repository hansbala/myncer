import { getPlaylistDetails } from "@/generated_grpc/myncer/datasource-DatasourceService_connectquery"
import type { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import { useQuery } from "@connectrpc/connect-query"

export const usePlaylist = ({ datasource, playlistId }: { datasource?: Datasource, playlistId?: string }) => {
  const { data: playlist, isLoading } = useQuery(
    getPlaylistDetails,
    {
      playlistId,
      datasource,
    },
    {
      // Only fetch playlists if a datasource and playlist id is provided.
      enabled: !!datasource && !!playlistId,
    },
  )
  return {
    playlist: playlist?.playlist,
    loading: isLoading,
  }
}
