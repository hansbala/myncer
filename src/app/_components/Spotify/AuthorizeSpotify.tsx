'use client'
import { api } from '~/trpc/react'

export default function AuthorizeSpotify() {
  const { isLoading, data: authUrl } =
    api.spotify.getAuthorizationUrl.useQuery()
  const { isLoading: isAuthenticatedLoading, data: isAuthenticated } =
    api.spotify.isAuthenticated.useQuery()

  if (isLoading || isAuthenticatedLoading) {
    return <div>Please wait...</div>
  }

  if (isAuthenticated) {
    return <div>✅ You have already authenticated Myncer for Spotify</div>
  }

  return (
    <button
      className="hover:underline"
      onClick={() => window.open(authUrl, '_blank')}
    >
      Authorize Myncer For Spotify
    </button>
  )
}
