'use client'

import { useState } from 'react'
import { type DATASOURCE } from '~/core/datasources'
import DatasourceDropdown from '../_components/DatasourceDropdown/DatasourceDropdown'
import PlaylistDropdown, {
  type Playlist,
} from '../_components/PlaylistDropdown/PlaylistDropdown'
import { MdOutlineSyncAlt } from 'react-icons/md'

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
            playlists={[]}
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
            playlists={[]}
            setPlaylist={setDestinationPlaylist}
            playlist={destinationPlaylist}
            disabled={!destinationDatasource}
            className="w-full"
          />
        </div>
      </div>
      <div className="divider"></div>
      <div className="flex flex-row items-center justify-center gap-4">
        <button className="btn btn-secondary" onClick={() => resetState()}>
          Reset
        </button>
        <button className="btn btn-primary">Create Sync</button>
      </div>
    </div>
  )
}
