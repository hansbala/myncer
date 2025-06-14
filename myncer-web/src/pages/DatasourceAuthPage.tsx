import { Datasource } from "@/generated_api/src"
import { useApiClient } from "@/hooks/useApiClient"
import { useEffect, useState } from "react"
import { useNavigate, useSearchParams, useParams } from "react-router-dom"

export const DatasourceAuthPage = () => {
  const apiClient = useApiClient()
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    const exchangeToken = async () => {
      const code = searchParams.get("code")
      const state = searchParams.get("state")

      if (!code) {
        setError("Missing required parameters.")
        return
      }

      try {
        await apiClient.exchangeOAuthCode({
          datasource: Datasource.Spotify,  // Hardcoded for now but later get proper typing.
          oAuthExchangeRequest: {
            code,
            state: state || undefined,
          },
        })
        navigate("/datasources", { replace: true })
      } catch (err) {
        console.error(err)
        setError("Authentication failed. Please try again.")
      }
    }

    exchangeToken()
  }, [apiClient, navigate, searchParams])

  return (
    <div className="flex h-screen items-center justify-center">
      {error ? (
        <div className="text-red-500 font-medium">{error}</div>
      ) : (
        <div className="text-muted-foreground">Linking your account…</div>
      )}
    </div>
  )
}
