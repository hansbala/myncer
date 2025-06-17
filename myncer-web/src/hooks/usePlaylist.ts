import type { Datasource, Playlist } from "@/generated_api/src";
import { useApiClient } from "./useApiClient";
import { useQuery } from "@tanstack/react-query";

export const usePlaylist = ({ datasource, playlistId }: { datasource?: Datasource, playlistId?: string }) => {
  const apiClient = useApiClient()
  const { data: playlist, isLoading } = useQuery<Playlist>({
    queryKey: ["playlists", "get", datasource, playlistId],
    queryFn: () => apiClient.getPlaylistDetails({ datasource: datasource!, playlistId: playlistId! }),
    // Only fetch playlists if a datasource and playlist id is provided.
    enabled: !!datasource && !!playlistId,
    retry: false,
  })
  return {
    playlist,
    loading: isLoading,
  }
}
