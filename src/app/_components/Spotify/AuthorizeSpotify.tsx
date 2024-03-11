'use client'
import { api } from '~/trpc/react'
import DatasourceConnection from '../DatasourceConnections/DatasourceConnection'
import { DATASOURCE_SCHEMAS } from '~/app/_core/clientSideDatasourceSchema'

export default function AuthorizeSpotify() {
  const { isLoading: isAuthenticationUrlLoading, data: authUrl } =
    api.spotify.getAuthorizationUrl.useQuery()
  const { isLoading: isAuthenticatedLoading, data: isAuthenticated } =
    api.spotify.isAuthenticated.useQuery()

  return (
    <DatasourceConnection
      datasourceSchema={DATASOURCE_SCHEMAS.SPOTIFY}
      isAuthenticatedLoading={isAuthenticatedLoading}
      isAuthenticationUrlLoading={isAuthenticationUrlLoading}
      authenticationUrl={authUrl}
      isAuthenticated={isAuthenticated}
    />
  )
}
