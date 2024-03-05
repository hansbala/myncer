"use client"
import { useRouter } from "next/navigation"
import { api } from "~/trpc/react"
import Button from "../_components/Button/Button"
import { useCallback, useEffect, useState } from "react"

export default function GoogleCallbackPage() {
  const router = useRouter()
  const [authorizationCode, setAuthorizationCode] = useState<string | null>()
  const updateAuthorizationCode = api.google.setAuthorizationCode.useMutation()

  useEffect(() => {
    const queryParamAuthCode = new URLSearchParams(window.location.search).get('code')
    setAuthorizationCode(queryParamAuthCode)
  }, [])

  const updateAuthCode = useCallback(() => {
    if (!authorizationCode) {
      return
    }
    updateAuthorizationCode.mutate({ authorizationCode }, {
      onSuccess: () => {
        void router.replace('/secrets')
      }
    })
  }, [authorizationCode, router, updateAuthorizationCode])

  return <>
    {
      !authorizationCode && (
        <div className="flex justify-center items-center h-screen">
          {
            <h1>Authorizing Google. Please do not close this page...</h1>
          }
        </div>
      )
    }

    {authorizationCode && (
      <div className="text-center p-10">
        <Button onClick={() => updateAuthCode()} className="">Authorized Google. Store auth token to finish setup.</Button>
      </div>)}
  </>
}