import { Datasource } from "@/generated_api/src"
import { useApiClient } from "@/hooks/useApiClient"
import { useEffect, useRef, useState } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"

export const DatasourceAuthPage = () => {
  const apiClient = useApiClient()
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const [error, setError] = useState<string | null>(null)
  // Used to make sure we only make one API call to the backend. In production, this is never an
  // issue but React strict mode always runs two useEffects so this guards against double calls.
  const didExchangeRef = useRef(false)

  useEffect(() => {
    const exchangeToken = async () => {
      if (didExchangeRef.current) return
      didExchangeRef.current = true
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
