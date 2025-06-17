import { toast } from "sonner"
import { useApiClient } from "./useApiClient"
import { useMutation, useQueryClient } from "@tanstack/react-query"

export const useDeleteSync = () => {
  const apiClient = useApiClient()
  const queryClient = useQueryClient()
  const { mutate: deleteSync, isPending: isDeleting } = useMutation({
    mutationKey: ["syncs", "delete"],
    mutationFn: async (syncId: string) => {
      return apiClient.deleteSync({ deleteSyncRequest: { syncId } })
    },
    onSuccess: () => {
      toast.success("Sync deleted!")
      queryClient.invalidateQueries({ queryKey: ["syncs"] })
    },
  })
  return {
    deleteSync,
    isDeleting,
  }
}
