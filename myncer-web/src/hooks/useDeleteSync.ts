import { deleteSync, listSyncs } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { createConnectQueryKey, useMutation } from "@connectrpc/connect-query"
import { useQueryClient } from "@tanstack/react-query"
import { toast } from "sonner"

export const useDeleteSync = () => {
  const queryClient = useQueryClient()
  const { mutate, isPending: isDeleting } = useMutation(deleteSync, {
    onSuccess: () => {
      toast.success("Sync deleted!")
      // Invalidate the syncs list so the new sync shows up in the UI
      queryClient.refetchQueries({
        queryKey: createConnectQueryKey({
          schema: listSyncs,
          cardinality: undefined,
        })
      })
    },
  })

  return {
    deleteSync: mutate,
    isDeleting,
  }
}
