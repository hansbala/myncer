import { listPlaylists } from "@/generated_grpc/myncer/datasource-DatasourceService_connectquery"
import type { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import { useQuery } from "@connectrpc/connect-query"

export const useListPlaylists = ({ datasource }: { datasource?: Datasource }) => {
  const { data: playlistsResponse, isLoading } = useQuery(
    listPlaylists,
    {
      datasource,
    },
    {
      // Only fetch playlists if a datasource is provided.
      enabled: !!datasource,
    },
  )
  return {
    playlists: playlistsResponse?.playlist ?? [],
    loading: isLoading,
  }
}
