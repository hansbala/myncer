import { runSync } from "@/generated_grpc/myncer/sync-SyncService_connectquery"
import { useMutation } from "@connectrpc/connect-query"
import { toast } from "sonner";

export const useRunSync = () => {
  const { mutateAsync, isPending: isRunningSync } = useMutation(runSync, {
    onSuccess: () => {
      toast.success("Sync started!");
    },
    onError: (error) => {
      toast.error(`Sync failed: ${error.message}`);
    },
  })
  return {
    runSync: mutateAsync,
    isRunningSync,
  }
}
