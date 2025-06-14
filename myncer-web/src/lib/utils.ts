import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

const spotifyScopes = [
  "user-read-email",
  "playlist-read-private",
  "playlist-modify-private",
  "playlist-modify-public",
  "user-library-read",
  "user-library-modify"
].join(" ")

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const getSpotifyAuthUrl = () => {
  const clientId = import.meta.env.VITE_SPOTIFY_CLIENT_ID
  const redirectUri = encodeURIComponent(import.meta.env.VITE_SPOTIFY_REDIRECT_URI)
  const scope = encodeURIComponent(spotifyScopes)
  const state = crypto.randomUUID() // CSRF protection.

  return `https://accounts.spotify.com/authorize?client_id=${clientId}&response_type=code&redirect_uri=${redirectUri}&scope=${scope}&state=${state}`
}
