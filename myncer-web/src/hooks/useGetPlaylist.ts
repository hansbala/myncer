import { getPlaylistDetails } from "@/generated_grpc/myncer/datasource-DatasourceService_connectquery"
import type { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import { useQuery } from "@connectrpc/connect-query"

export const useGetPlaylist = ({ datasource, playlistId }: { datasource?: Datasource, playlistId?: string }) => {
  const { data: response, isLoading, isError } = useQuery(
    getPlaylistDetails,
    {
      datasource,
      playlistId,
    },
    {
      // Only fetch playlists if a datasource and playlist id is provided.
      enabled: !!datasource && !!playlistId,
    },
  )

  return {
    playlist: isError ? undefined : response?.playlist,
    loading: isLoading,
  }
}
