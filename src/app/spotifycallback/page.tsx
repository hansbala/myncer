"use client"
import { useRouter } from "next/navigation"
import { api } from "~/trpc/react"
import Button from "../_components/Button/Button"
import { useCallback, useEffect, useState } from "react"

export default function SpotifyCallbackPage() {
  const router = useRouter()
  const [authCode, setAuthCode] = useState<string | null>()
  const updateAuthorizationCode = api.spotify.setAuthorizationCode.useMutation()

  useEffect(() => {
    const queryParamAuthCode = new URLSearchParams(window.location.search).get('code')
    setAuthCode(queryParamAuthCode)
  }, [])

  const updateAuthCode = useCallback(() => {
    if (!authCode) {
      return
    }
    updateAuthorizationCode.mutate({ authCode }, {
      onSuccess: () => {
        void router.replace('/secrets')
      }
    })
  }, [authCode, router, updateAuthorizationCode])

  return <>
    {
      !authCode && (
        <div className="flex justify-center items-center h-screen">
          {
            <h1>Auhorizing Spotify. Please do not close this page...</h1>
          }
        </div>
      )
    }

    {authCode && (
      <div className="text-center p-10">
        <Button onClick={() => updateAuthCode()} className="">Authorized Spotify. Store auth token to finish setup.</Button>
      </div>)}
  </>
}