'use client'
import { DATASOURCE_SCHEMAS } from '~/app/_core/clientSideDatasourceSchema'
import { api } from '~/trpc/react'
import DatasourceConnection from '../DatasourceConnections/DatasourceConnection'

export default function AuthorizeGoogle() {
  const { isLoading: isAuthenticationUrlLoading, data: authUrl } =
    api.google.getAuthorizationUrl.useQuery()
  const { isLoading: isAuthenticatedLoading, data: isAuthenticated } =
    api.google.isAuthenticated.useQuery()

  return (
    <DatasourceConnection
      datasourceSchema={DATASOURCE_SCHEMAS.YOUTUBE}
      isAuthenticatedLoading={isAuthenticatedLoading}
      isAuthenticationUrlLoading={isAuthenticationUrlLoading}
      authenticationUrl={authUrl}
      isAuthenticated={isAuthenticated}
    />
  )
}
