import { listSyncs } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { useQuery } from "@connectrpc/connect-query"

export const useSyncs = () => {
  const { data: listSyncsResponse, isLoading } = useQuery(listSyncs)
  return {
    syncs: listSyncsResponse?.syncs ?? [],
    loading: isLoading,
  }
}
