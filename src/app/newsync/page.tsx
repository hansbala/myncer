'use client'

import { useCallback, useState } from 'react'
import { DATASOURCE } from '~/core/datasources'
import DatasourceDropdown from '../_components/DatasourceDropdown/DatasourceDropdown'
import PlaylistDropdown, {
  type Playlist,
} from '../_components/PlaylistDropdown/PlaylistDropdown'
import { MdOutlineSyncAlt } from 'react-icons/md'
import { api } from '~/trpc/react'
import { type youtube_v3 } from 'googleapis'
import { SYNC_FREQUENCY } from '~/core/syncFrequency'

export default function NewSyncPage() {
  const [sourceDatasource, setSourceDatasource] = useState<DATASOURCE>()
  const [destinationDatasource, setDestinationDatasource] =
    useState<DATASOURCE>()
  const [sourcePlaylist, setSourcePlaylist] = useState<Playlist>()
  const [destinationPlaylist, setDestinationPlaylist] = useState<Playlist>()
  const [syncFrequency, setSyncFrequency] = useState<SYNC_FREQUENCY>(
    SYNC_FREQUENCY.DAILY
  )
  const createSyncOnServer = api.syncs.createNewSync.useMutation()

  // TODO(Hans): Abstract away datasource specific logic into a hook useAllPlaylistData()
  const { data: spotifyPlaylistData } =
    api.spotify.getCurrentUserPlaylists.useQuery()
  const { data: youtubePlaylistData } =
    api.google.getCurrentUserPlaylists.useQuery()

  const getPlaylistDataByDatasource = useCallback(
    (datasource: DATASOURCE) => {
      switch (datasource) {
        case DATASOURCE.SPOTIFY:
          return spotifyPlaylistData?.items
        case DATASOURCE.YOUTUBE:
          return youtubePlaylistData
        default:
          return []
      }
    },
    [spotifyPlaylistData?.items, youtubePlaylistData]
  )

  const getPlaylists = useCallback(
    (type: 'source' | 'destination') => {
      const datasource =
        type === 'source' ? sourceDatasource : destinationDatasource
      if (!datasource) return []
      const playlists = getPlaylistDataByDatasource(datasource)
      if (!playlists) return []
      // TODO(Hans): clean this up
      switch (datasource) {
        case DATASOURCE.SPOTIFY:
          return (playlists as SpotifyApi.PlaylistObjectSimplified[]).map(
            (playlist) => ({
              id: playlist.id,
              name: playlist.name,
            })
          ) as Playlist[]
      }
      return (playlists as youtube_v3.Schema$Playlist[]).map((playlist) => ({
        id: playlist.id,
        name: playlist.snippet?.title ?? '',
      })) as Playlist[]
    },
    [sourceDatasource, destinationDatasource, getPlaylistDataByDatasource]
  )

  const resetState = () => {
    setSourceDatasource(undefined)
    setDestinationDatasource(undefined)
    setSourcePlaylist(undefined)
    setDestinationPlaylist(undefined)
  }

  const createSync = async () => {
    if (
      !sourcePlaylist ||
      !destinationPlaylist ||
      !sourceDatasource ||
      !destinationDatasource
    ) {
      return
    }
    const input = {
      sourcePlaylist: {
        datasource: sourceDatasource,
        playlistId: sourcePlaylist.id,
      },
      destinationPlaylist: {
        datasource: destinationDatasource,
        playlistId: destinationPlaylist.id,
      },
      syncFrequency,
      useWebhooks: false,
    }
    await createSyncOnServer.mutateAsync(input)
    resetState()
  }

  return (
    <div className="flex w-1/4 flex-col gap-10 pt-10">
      <div className="flex flex-row items-center justify-between">
        <div className="flex flex-col items-center justify-center">
          <h3 className="mb-5 text-xl font-bold">Source</h3>
          <DatasourceDropdown
            datasource={sourceDatasource}
            setDatasource={setSourceDatasource}
            className="mb-2 w-full"
          />
          <PlaylistDropdown
            playlists={getPlaylists('source')}
            setPlaylist={setSourcePlaylist}
            playlist={sourcePlaylist}
            disabled={!sourceDatasource}
            className="w-full"
          />
        </div>
        <div className="divider divider-horizontal">
          <MdOutlineSyncAlt size={128} />
        </div>
        <div className="flex flex-col items-center justify-center">
          <h3 className="mb-5 text-xl font-bold">Destination</h3>
          <DatasourceDropdown
            datasource={destinationDatasource}
            setDatasource={setDestinationDatasource}
            className="mb-2 w-full"
          />
          <PlaylistDropdown
            playlists={getPlaylists('destination')}
            setPlaylist={setDestinationPlaylist}
            playlist={destinationPlaylist}
            disabled={!destinationDatasource}
            className="w-full"
          />
        </div>
      </div>
      <div>
        <h3 className="mb-5 text-center text-xl font-bold">Sync Frequency</h3>
        <div className="flex flex-row items-center justify-center gap-4">
          <input
            type="radio"
            id="daily"
            name="syncFrequency"
            value={SYNC_FREQUENCY.DAILY}
            checked={syncFrequency === SYNC_FREQUENCY.DAILY}
            onChange={(e) => setSyncFrequency(e.target.value as SYNC_FREQUENCY)}
          />
          <label htmlFor="daily">Daily</label>
          <input
            type="radio"
            id="weekly"
            name="syncFrequency"
            value={SYNC_FREQUENCY.WEEKLY}
            checked={syncFrequency === SYNC_FREQUENCY.WEEKLY}
            onChange={(e) => setSyncFrequency(e.target.value as SYNC_FREQUENCY)}
          />
          <label htmlFor="weekly">Weekly</label>
          <input
            type="radio"
            id="monthly"
            name="syncFrequency"
            value={SYNC_FREQUENCY.MONTHLY}
            checked={syncFrequency === SYNC_FREQUENCY.MONTHLY}
            onChange={(e) => setSyncFrequency(e.target.value as SYNC_FREQUENCY)}
          />
          <label htmlFor="monthly">Monthly</label>
        </div>
      </div>
      <div className="divider"></div>
      <div className="flex flex-row items-center justify-center gap-4">
        <button className="btn btn-secondary" onClick={() => resetState()}>
          Reset
        </button>
        <button className="btn btn-primary" onClick={() => createSync()}>
          Create Sync
        </button>
      </div>
    </div>
  )
}
