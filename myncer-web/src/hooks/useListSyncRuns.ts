import { listSyncRuns } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { protoTimestampToDate } from "@/lib/utils"
import { useQuery } from "@connectrpc/connect-query"

export const useListSyncRuns = () => {
  const { data: syncRunsResponse, isLoading } = useQuery(listSyncRuns)
  // Sort by createdAt in descending order.
  const sortedSyncRuns = syncRunsResponse?.syncRuns?.sort((a, b) => {
    if (!a.createdAt || !b.createdAt) return 0
    return new Date(protoTimestampToDate(b.createdAt)).getTime() - 
      new Date(protoTimestampToDate(a.createdAt)).getTime()
  })
  return {
    syncRuns: sortedSyncRuns ?? [],
    loading: isLoading,
  }
}
