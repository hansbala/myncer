import { google } from "googleapis"
import { env } from "~/env"

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