import type { Datasource, ListDatasourcePlaylistsResponse } from "@/generated_api/src";
import { useApiClient } from "./useApiClient";
import { useQuery } from "@tanstack/react-query";

export const usePlaylists = ({ datasource }: { datasource?: Datasource }) => {
  const apiClient = useApiClient()
  const { data: playlistsResponse, isLoading } = useQuery<ListDatasourcePlaylistsResponse>({
    queryKey: ["playlists", "list", datasource],
    queryFn: () => apiClient.listDatasourcePlaylists({ datasource: datasource! }),
    // Only fetch playlists if a datasource is provided.
    enabled: !!datasource,
    retry: false,
  })
  return {
    playlists: playlistsResponse ? playlistsResponse.playlists || [] : [],
    loading: isLoading,
  }
}
