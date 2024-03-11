'use client'
import { useRouter } from 'next/navigation'
import { api } from '~/trpc/react'
import Button from '../_components/Button/Button'
import { useCallback, useEffect, useState } from 'react'

export default function SpotifyCallbackPage() {
  const router = useRouter()
  const [authCode, setAuthCode] = useState<string | null>()
  const updateAuthorizationCode = api.spotify.setAuthorizationCode.useMutation()

  useEffect(() => {
    const queryParamAuthCode = new URLSearchParams(window.location.search).get(
      'code'
    )
    setAuthCode(queryParamAuthCode)
  }, [])

  const updateAuthCode = useCallback(() => {
    if (!authCode) {
      return
    }
    updateAuthorizationCode.mutate(
      { authCode },
      {
        onSuccess: () => {
          void router.replace('/secrets')
        },
      }
    )
  }, [authCode, router, updateAuthorizationCode])

  return (
    <>
      {!authCode && (
        <div className="flex h-screen items-center justify-center">
          {<h1>Authorizing Spotify. Please do not close this page...</h1>}
        </div>
      )}

      {authCode && (
        <div className="p-10 text-center">
          <Button onClick={() => updateAuthCode()} className="">
            Authorized Spotify. Store auth token to finish setup.
          </Button>
        </div>
      )}
    </>
  )
}
