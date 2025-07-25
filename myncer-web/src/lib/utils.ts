import { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import type { Timestamp } from "@bufbuild/protobuf/wkt"
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

const youtubeScopes = [
  "https://www.googleapis.com/auth/youtube"
].join(" ")

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const getSpotifyAuthUrl = () => {
  const clientId = import.meta.env.VITE_SPOTIFY_CLIENT_ID
  const redirectUri = encodeURIComponent(import.meta.env.VITE_SPOTIFY_REDIRECT_URI)
  const scope = encodeURIComponent(spotifyScopes)
  const state = crypto.randomUUID() // CSRF protection.

  return `https://accounts.spotify.com/authorize?client_id=${clientId}&response_type=code&redirect_uri=${redirectUri}&scope=${scope}&state=${state}&prompt=consent`
}

export const getYoutubeAuthUrl = () => {
  const clientId = import.meta.env.VITE_YOUTUBE_CLIENT_ID
  const redirectUri = encodeURIComponent(import.meta.env.VITE_YOUTUBE_REDIRECT_URI)
  const scope = encodeURIComponent(youtubeScopes)
  const state = crypto.randomUUID() // CSRF protection

  return `https://accounts.google.com/o/oauth2/v2/auth?client_id=${clientId}&redirect_uri=${redirectUri}&response_type=code&scope=${scope}&state=${state}&access_type=offline&prompt=consent`
}

export const getDatasourceLabel = (datasource: Datasource) => {
  switch (datasource) {
    case Datasource.SPOTIFY:
      return "Spotify"
    case Datasource.YOUTUBE:
      return "YouTube"
    default:
      return "Unknown Datasource"
  }
}

export const protoTimestampToDate = (ts: Timestamp): Date => {
  const millis = Number(ts.seconds) * 1000 + Math.floor((ts.nanos || 0) / 1_000_000)
  return new Date(millis)
}

