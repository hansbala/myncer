'use client'
import { api } from '~/trpc/react'

export default function AuthorizeGoogle() {
  const { isLoading, data: authUrl } = api.google.getAuthorizationUrl.useQuery()
  const { isLoading: isAuthenticatedLoading, data: isAuthenticated } =
    api.google.isAuthenticated.useQuery()

  if (isLoading || isAuthenticatedLoading) {
    return <div>Please wait...</div>
  }

  if (isAuthenticated) {
    return <div>✅ You have already authenticated Myncer for Google</div>
  }

  return (
    <button
      className="hover:underline"
      onClick={() => window.open(authUrl, '_blank')}
    >
      Authorize Myncer For Google
    </button>
  )
}
