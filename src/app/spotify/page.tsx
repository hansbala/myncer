'use client'

import { useMemo } from 'react'
import { api } from '~/trpc/react'

export default function SpotifyPage() {
  const { data, isLoading } = api.spotify.getCurrentUserPlaylists.useQuery()

  const playlists = useMemo(() => {
    return data?.items.map((playlist) => {
      return {
        name: playlist.name,
        url: playlist.uri,
        image: playlist.images[0],
      }
    })
  }, [data])

  if (isLoading) {
    return <div>Loading....</div>
  }

  return (
    <>
      <h1>Current Playlists</h1>
      <div className="flex flex-col gap-5 overflow-y-visible">
        {playlists?.map((playlist) => (
          <div key={playlist.url} className="flex flex-row items-center gap-3">
            {/* eslint-disable-next-line @next/next/no-img-element */}
            <img
              alt="Playlist image"
              src={playlist.image?.url ?? '/test.img'}
              width={100}
              height={100}
            />
            <h1>{playlist.name}</h1>
            <p>{playlist.url}</p>
          </div>
        ))}
      </div>
    </>
  )
}
