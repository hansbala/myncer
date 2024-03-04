import SpotifyWebApi from 'spotify-web-api-node'
import { getRandomString } from './crypto'
import { db } from '../db'
import { env } from '~/env'

const MYNCER_CLIENT_ID = env.MYNCER_CLIENT_ID
const MYNCER_REDIRECT_URL = env.MYNCER_REDIRECT_URL
const MYNCER_CLIENT_SECRET = env.MYNCER_CLIENT_SECRET

const MYNCER_REQUIRED_SCOPES = ['user-read-playback-state',
  'user-modify-playback-state',
  'user-read-currently-playing',
  'app-remote-control',
  'streaming',
  'playlist-read-private',
  'playlist-read-collaborative',
  'playlist-modify-private',
  'playlist-modify-public',
  'user-follow-modify',
  'user-follow-read',
  'user-read-playback-position',
  'user-top-read',
  'user-read-recently-played',
  'user-library-modify',
  'user-library-read',
  'user-read-email',
  'user-read-private']

export const authUserForSpotifyFirstTime = async (authCode: string): Promise<{ accessToken: string, refreshToken: string }> => {
  const spotifyApi = new SpotifyWebApi({
    redirectUri: MYNCER_REDIRECT_URL,
    clientId: MYNCER_CLIENT_ID,
    clientSecret: MYNCER_CLIENT_SECRET,
  })
  const { body: { access_token: accessToken, refresh_token: refreshToken } } = await spotifyApi.authorizationCodeGrant(authCode)
  return {
    accessToken,
    refreshToken
  }
}

const getAccessAndRefreshTokens = async (myncerUserId: string): Promise<{ accessToken: string, refreshToken: string }> => {
  const spotifyApi = new SpotifyWebApi({
    redirectUri: MYNCER_REDIRECT_URL,
    clientId: MYNCER_CLIENT_ID,
    clientSecret: MYNCER_CLIENT_SECRET,
  })

  // fetch if user has exisiting accessToken from db
  const findUserResult = await db.spotifyApiKey.findUnique({
    where: {
      userId: myncerUserId,
    },
    select: {
      accessToken: true,
      refreshToken: true
    }
  })

  if (findUserResult == null) {
    throw new Error('Could not find user')
  }

  spotifyApi.setAccessToken(findUserResult.accessToken)
  spotifyApi.setRefreshToken(findUserResult.refreshToken)
  // refresh the access token
  const res = await spotifyApi.refreshAccessToken()

  // push tokens to database
  const { accessToken, refreshToken } = await db.spotifyApiKey.update({
    where: {
      userId: myncerUserId
    },
    data: {
      accessToken: res.body.access_token,
      refreshToken: res.body.refresh_token
    }
  })

  return {
    accessToken,
    refreshToken
  }
}

export const getCurrentUserPlaylists = async (myncerUserId: string) => {
  const { accessToken, refreshToken } = await getAccessAndRefreshTokens(myncerUserId)
  const spotifyApi = new SpotifyWebApi({
    redirectUri: MYNCER_REDIRECT_URL,
    clientId: MYNCER_CLIENT_ID,
    clientSecret: MYNCER_CLIENT_SECRET,
  })
  spotifyApi.setAccessToken(accessToken)
  spotifyApi.setRefreshToken(refreshToken)

  const playlists = await spotifyApi.getUserPlaylists()
  return playlists.body
}

export const createSpotifyAuthUrl = (): string => {
  const state = getRandomString(16)
  const spotifyApi = new SpotifyWebApi({
    redirectUri: MYNCER_REDIRECT_URL,
    clientId: MYNCER_CLIENT_ID,
  })
  return spotifyApi.createAuthorizeURL(MYNCER_REQUIRED_SCOPES, state)
}