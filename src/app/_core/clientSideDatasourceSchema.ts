import { FaSpotify, FaYoutube } from 'react-icons/fa'
import { type IconType } from 'react-icons'
import { DATASOURCE } from '~/core/datasources'

export interface DATASOURCE_SCHEMA {
  datasource: DATASOURCE
  name: string
  description: string
  clientIcon: IconType
}

export const DATASOURCE_SCHEMAS: Record<DATASOURCE, DATASOURCE_SCHEMA> = {
  [DATASOURCE.SPOTIFY]: {
    datasource: DATASOURCE.SPOTIFY,
    name: 'Spotify',
    description: 'Connect to Spotify to sync your playlists with Myncer',
    clientIcon: FaSpotify,
  },
  [DATASOURCE.YOUTUBE]: {
    datasource: DATASOURCE.YOUTUBE,
    name: 'YouTube',
    description: 'Connect to YouTube to sync your playlists with Myncer',
    clientIcon: FaYoutube,
  },
}
