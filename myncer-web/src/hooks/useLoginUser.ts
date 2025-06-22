import { getCurrentUser, loginUser } from "@/generated_grpc/myncer/user-UserService_connectquery"
import { createConnectQueryKey, useMutation } from "@connectrpc/connect-query"
import { useQueryClient } from "@tanstack/react-query"
import { toast } from "sonner"

export const useLoginUser = () => {
  const queryClient = useQueryClient()
  return useMutation(loginUser, {
    onSuccess: () => {
      toast.success("Hey there!")
      const queryKey = createConnectQueryKey({
        schema: getCurrentUser,
        cardinality: undefined,
      })
      queryClient.refetchQueries({
        queryKey,
      })
    },
    onError: (error) => {
      toast.error(`Failed to login: ${error.rawMessage}`)
    }
  })
}
