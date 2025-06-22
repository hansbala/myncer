import * as DatasourceService from "@/generated_grpc/myncer/datasource-DatasourceService_connectquery"
import type { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import { createConnectQueryKey, useMutation, useTransport } from "@connectrpc/connect-query"
import { useQueryClient } from "@tanstack/react-query"
import { useEffect, useRef } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { toast } from "sonner"

interface DatasourceAuthPageProps {
  datasource: Datasource
}
export const DatasourceAuthPage = ({ datasource }: DatasourceAuthPageProps) => {
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  // Used to make sure we only make one API call to the backend. In production, this is never an
  // issue but React strict mode always runs two useEffects so this guards against double calls.
  const didExchangeRef = useRef(false)

  const queryClient = useQueryClient()
  const transport = useTransport()
  const { mutateAsync: exchangeOAuthCode } = useMutation(DatasourceService.exchangeOAuthCode, {
    onSuccess: () => {
      toast.success("OAuth exchange successful!")
      const key = createConnectQueryKey({
        schema: DatasourceService.listDatasources,
        transport,
        input: {},
        cardinality: "finite",
      })
      queryClient.setQueryData(key, undefined)
      queryClient.invalidateQueries({ queryKey: key })
    },
    onError: (err) => {
      toast.error(`OAuth Exchange failed: ${err.message}`)
    },
  })

  useEffect(() => {
    const exchangeToken = async () => {
      if (didExchangeRef.current) return
      didExchangeRef.current = true
      const code = searchParams.get("code")
      const state = searchParams.get("state")
      if (!code) {
        toast.error("Missing OAuth code in URL parameters.")
        return
      }
      try {
        await exchangeOAuthCode({
          datasource,
          code,
          csrfToken: state || undefined,
        })
        navigate("/datasources", { replace: true })
      } catch (err) {
        console.error(err)
      }
    }

    exchangeToken()
  }, [navigate, searchParams])

  return (
    <div className="flex h-screen items-center justify-center">
      <div className="text-muted-foreground">Linking your account…</div>
    </div>
  )
}
