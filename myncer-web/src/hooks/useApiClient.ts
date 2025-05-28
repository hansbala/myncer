import { useMemo } from "react"
import { Configuration, DefaultApi } from "../generated_api/src"

const BASE_PATH = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1'

export const useApiClient = () => {
  return useMemo(() => {
    const config = new Configuration({ basePath: BASE_PATH, credentials: 'include' })
    return new DefaultApi(config)
  }, [])
}
