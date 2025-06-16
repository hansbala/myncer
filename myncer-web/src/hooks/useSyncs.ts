import { useQuery } from "@tanstack/react-query"
import { useApiClient } from "./useApiClient"

export const useSyncs = () => {
  const apiClient = useApiClient()
  const { data: listSyncsResponse, isLoading } = useQuery({
    queryKey: ["syncs", "list"],
    queryFn: () => apiClient.listSyncs(),
    retry: false,
  })
  return {
    syncs: listSyncsResponse ? listSyncsResponse.syncs || [] : [],
    loading: isLoading,
  }
}
