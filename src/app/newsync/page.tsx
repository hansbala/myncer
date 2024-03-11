'use client'

import { useState } from 'react'
import { type DATASOURCE } from '~/core/datasources'
import DatasourceDropdown from '../_components/DatasourceDropdown/DatasourceDropdown'
import PlaylistDropdown, {
  type Playlist,
} from '../_components/PlaylistDropdown/PlaylistDropdown'

export default function NewSyncPage() {
  const [sourceDatasource, setSourceDatasource] = useState<DATASOURCE>()
  const [destinationDatasource, setDestinationDatasource] =
    useState<DATASOURCE>()
  const [sourcePlaylist, setSourcePlaylist] = useState<Playlist>()
  const [destinationPlaylist, setDestinationPlaylist] = useState<Playlist>()

  const resetState = () => {
    setSourceDatasource(undefined)
    setDestinationDatasource(undefined)
    setSourcePlaylist(undefined)
    setDestinationPlaylist(undefined)
  }

  return (
    <div className="flex flex-col">
      <h2>Create a new sync here</h2>
      <div className="flex flex-row items-center">
        <div className="pr-10">Source:</div>
        <DatasourceDropdown
          datasource={sourceDatasource}
          setDatasource={setSourceDatasource}
        />
      </div>
      <div className="flex flex-row items-center">
        <div className="pr-10">Destination:</div>
        <DatasourceDropdown
          datasource={destinationDatasource}
          setDatasource={setDestinationDatasource}
        />
      </div>
      <div className="flex flex-row items-center">
        <div className="pr-10">Source:</div>
        <PlaylistDropdown
          playlists={[]}
          setPlaylist={setSourcePlaylist}
          playlist={sourcePlaylist}
        />
      </div>
      <div className="flex flex-row items-center">
        <div className="pr-10">Destination:</div>
        <PlaylistDropdown
          playlists={[]}
          setPlaylist={setDestinationPlaylist}
          playlist={destinationPlaylist}
        />
      </div>

      <div className="flex flex-row items-center justify-center gap-4">
        <button className="btn btn-primary">Create Sync</button>
        <button className="btn btn-ghost" onClick={() => resetState()}>
          Reset
        </button>
      </div>
    </div>
  )
}
