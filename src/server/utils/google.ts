import { google } from "googleapis"
import { env } from "~/env"
import { db } from "../db"

const youtubePlaylistManagementScopes = [
  'https://www.googleapis.com/auth/youtube.readonly',
  'https://www.googleapis.com/auth/youtube'
]

/**
 * Generates a auth URL that user will navigate to to authorize the app.
 * This one is specifically for granting myncer access to manage youtube playlists
 * for syncing purposes
 */
export const getGoogleAuthorizationUrl = () => {
  const oauth2Client = new google.auth.OAuth2({
    clientId: env.GOOGLE_MYNCER_CLIENT_ID,
    clientSecret: env.GOOGLE_MYNCER_CLIENT_SECRET,
    redirectUri: env.GOOGLE_MYNCER_REDIRECT_URL
  })

  return oauth2Client.generateAuthUrl({
    // default is 'online' but 'offline' gets refresh token (necessary for long-lived apps like myncer)
    access_type: 'offline',
    scope: youtubePlaylistManagementScopes
  })
}


export const authUserForGoogleFirstTime = async (authCode: string): Promise<{ accessToken: string, refreshToken: string }> => {
  const oauth2Client = new google.auth.OAuth2({
    clientId: env.GOOGLE_MYNCER_CLIENT_ID,
    clientSecret: env.GOOGLE_MYNCER_CLIENT_SECRET,
    redirectUri: env.GOOGLE_MYNCER_REDIRECT_URL
  })

  const { tokens: { refresh_token: refreshToken, access_token: accessToken } } = await oauth2Client.getToken(authCode)
  console.log('access token', accessToken)
  console.log('refresh token', refreshToken)
  if (!refreshToken || !accessToken) {
    throw new Error('Failed to get either access token or refresh token for first time user auth')
  }

  return {
    refreshToken,
    accessToken
  }
}

export const getGoogleKeyForUser = async (myncerUserId: string) => {
  const foundKey = await db.googleKey.findUnique({
    where: {
      userId: myncerUserId
    }
  })
  if (!foundKey) {
    throw new Error('No google key found for user')
  }
  return {
    accessToken: foundKey.accessCode,
    refreshToken: foundKey.refreshToken
  }
}

export const getCurrentUserPlaylists = async (myncerUserId: string) => {
  const { accessToken, refreshToken } = await getGoogleKeyForUser(myncerUserId)
  const oauth2Client = new google.auth.OAuth2({
    clientId: env.GOOGLE_MYNCER_CLIENT_ID,
    clientSecret: env.GOOGLE_MYNCER_CLIENT_SECRET,
    redirectUri: env.GOOGLE_MYNCER_REDIRECT_URL
  })
  oauth2Client.setCredentials({
    access_token: accessToken,
    refresh_token: refreshToken
  })
  // using googleapis to get the youtube client
  const youtube = google.youtube({
    version: 'v3',
    auth: oauth2Client
  })
  const { data } = await youtube.playlists.list({
    part: ['snippet'],
    mine: true
  })
  return data.items ?? []
}
