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
        <div>
          <DatasourceDropdown
            datasource={sourceDatasource}
            setDatasource={setSourceDatasource}
          />
          <PlaylistDropdown
            playlists={[]}
            setPlaylist={setSourcePlaylist}
            playlist={sourcePlaylist}
            disabled={!sourceDatasource}
          />
        </div>
        <MdOutlineSyncAlt size={52} />
        <div>
          <DatasourceDropdown
            datasource={destinationDatasource}
            setDatasource={setDestinationDatasource}
          />
          <PlaylistDropdown
            playlists={[]}
            setPlaylist={setDestinationPlaylist}
            playlist={destinationPlaylist}
            disabled={!destinationDatasource}
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
