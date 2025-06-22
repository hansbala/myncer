import { createSync, listSyncs } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { createConnectQueryKey, useMutation } from "@connectrpc/connect-query"
import { useQueryClient } from "@tanstack/react-query"
import { toast } from "sonner"

export const useCreateSync = () => {
  const queryClient = useQueryClient()
  return useMutation(createSync, {
    onSuccess: () => {
      toast.success("Sync created!")
      // Invalidate the syncs list so the new sync shows up in the UI
      queryClient.invalidateQueries({
        queryKey: createConnectQueryKey({
          schema: listSyncs,
          cardinality: undefined,
        })
      })
    },
  })
}
