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