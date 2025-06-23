import { getSync } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { useQuery } from "@connectrpc/connect-query"

export const useSync = (syncId: string | undefined) => {
  const { data, isLoading } = useQuery(
    getSync,
    { syncId: syncId! },
    { enabled: !!syncId },
  )
  return {
    sync: data?.sync,
    loading: isLoading,
  }
}
