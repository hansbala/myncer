"use client"

import { useMemo } from "react"
import { api } from "~/trpc/react"

export default function SpotifyPage() {
  const { data, isLoading } = api.spotify.getCurrentUserPlaylists.useQuery()

  const playlists = useMemo(() => {
    return data?.items.map((playlist) => {
      return playlist.name
    })
  }, [data])

  if (isLoading) {
    return <div>Loading....</div>
  }

  return (
    <>
      <h1>Current Playlists</h1>
      <div>
        {JSON.stringify(playlists, null, 2)}
      </div>
    </>
  )
}