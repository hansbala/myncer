import { useQuery, useQueryClient } from "@tanstack/react-query"
import { useApiClient } from "./useApiClient"
import type { User } from "../generated_api/src/models"

export const useAuth = () => {
  const apiClient = useApiClient()
  const queryClient = useQueryClient()

  const { data: user, isLoading, isError } = useQuery<User>({
    queryKey: ["auth", "me"],
    queryFn: () => apiClient.getCurrentUser(),
    retry: false,
  })

  const logout = async () => {
    try {
      await apiClient.logoutUser()
    } catch (err) {
      console.error("Logout failed:", err)
    } finally {
      queryClient.setQueryData(["auth", "me"], null)
      queryClient.invalidateQueries({ queryKey: ["auth", "me"] })
    }
  }

  return {
    user,
    isAuthenticated: !!user,
    loading: isLoading,
    error: isError,
    logout,
  }
}
