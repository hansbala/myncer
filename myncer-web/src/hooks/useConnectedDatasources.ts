import type { ListDatasourcesResponse } from "@/generated_api/src"
import { useApiClient } from "./useApiClient"
import { useQuery } from "@tanstack/react-query"

export const useConnectedDatasources = () => {
  const apiClient = useApiClient()
  const { data: listDatasourcesResponse, isLoading } = useQuery<ListDatasourcesResponse>({
    queryKey: ["datasources", "list"],
    queryFn: () => apiClient.listConnectedDatasources(),
    retry: false,
  })
  return {
    datasources: listDatasourcesResponse ? listDatasourcesResponse.connectedDatasources || [] : [],
    loading: isLoading,
  }
}
