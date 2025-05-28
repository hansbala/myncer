import { useQuery } from "@tanstack/react-query"
import { useApiClient } from "./useApiClient"
import type { User } from "../generated_api/src/models"

export const useAuth = () => {
  const apiClient = useApiClient()

  const { data: user, isLoading, isError } = useQuery<User>({
    queryKey: ["auth", "me"],
    queryFn: () => apiClient.getCurrentUser(),
    retry: false,
  })

  return {
    user,
    isAuthenticated: !!user,
    loading: isLoading,
    error: isError,
  }
}
