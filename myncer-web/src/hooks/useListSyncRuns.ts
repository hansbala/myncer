import { listSyncRuns } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { useQuery } from "@connectrpc/connect-query"

export const useListSyncRuns = () => {
  const { data: syncRunsResponse, isLoading } = useQuery(listSyncRuns)
  return {
    syncRuns: syncRunsResponse?.syncRuns ?? [],
    loading: isLoading,
  }
}
