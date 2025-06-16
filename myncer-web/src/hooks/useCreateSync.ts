import { useMutation, useQueryClient } from "@tanstack/react-query"
import { useApiClient } from "./useApiClient"
import type { CreateSyncRequest } from "@/generated_api/src"

export const useCreateSync = () => {
  const apiClient = useApiClient()
  const queryClient = useQueryClient()

  return useMutation({
    mutationFn: async (payload: CreateSyncRequest): Promise<void> => {
      return apiClient.createSync({
        createSyncRequest: payload,
      })
    },
    onSuccess: () => {
      // Invalidate the syncs list so the new sync shows up in the UI
      queryClient.invalidateQueries({ queryKey: ["syncs", "list"] })
    },
  })
}

