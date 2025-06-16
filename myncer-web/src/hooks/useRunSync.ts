import { useMutation } from "@tanstack/react-query";
import { useApiClient } from "./useApiClient";
import { toast } from "sonner"

export const useRunSync = () => {
  const apiClient = useApiClient();
  const { mutate: runSync, isPending: isRunningSync } = useMutation({
    mutationFn: (id: string) => apiClient.runSync({
      runSyncRequest: {
        syncId: id,
      }
    }),
    onSuccess: () => {
      toast.success("Sync started!");
    },
    onError: (error) => {
      toast.error(`Sync failed: ${error.message}`);
    },
  })
  return {
    runSync,
    isRunningSync,
  }
}
