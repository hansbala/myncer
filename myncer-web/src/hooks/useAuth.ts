import { getCurrentUser, logoutUser } from "@/generated_grpc/myncer/user-UserService_connectquery"
import { createConnectQueryKey, useMutation, useQuery, useTransport } from "@connectrpc/connect-query"
import { useQueryClient } from "@tanstack/react-query"
import { toast } from "sonner"

export const useAuth = () => {
  const queryClient = useQueryClient()
  const transport = useTransport()

  const {
    data: currentUserResponse,
    isLoading,
    isError,
  } = useQuery(getCurrentUser, {}, { retry: false, staleTime: 0 })

  const logout = useMutation(logoutUser, {
    onSuccess: () => {
      toast.success("Logged out!")
      const key = createConnectQueryKey({
        schema: getCurrentUser,
        transport,
        input: {},
        cardinality: "finite",
      })
      // Remove the stale user from cache
      queryClient.setQueryData(key, undefined)
      // Optional: Force a refetch if desired (will error 403 and be ignored)
      queryClient.invalidateQueries({ queryKey: key })
    },
    onError: (err) => {
      toast.error(`Logout failed: ${err.message}`)
    },
  })

  return {
    user: isError ? undefined : currentUserResponse?.user,
    isAuthenticated: !isError && !!currentUserResponse?.user,
    loading: isLoading,
    error: isError,
    logout: logout,
  }
}
