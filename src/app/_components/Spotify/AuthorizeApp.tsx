"use client"
import { api } from "~/trpc/react"

export default function AuthorizeApp() {
  const { isLoading, data: authUrl } = api.spotify.getAuthorizationUrl.useQuery()


  return (
    <>
      {
        isLoading ? <div>Please wait ...</div>
          :
          <button className="hover:underline" onClick={() => window.open(authUrl, "_blank")}>Authorize Myncer</button>
      }
    </>
  )
}
