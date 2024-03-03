"use client"

import { api } from "~/trpc/react"

export default function SpotifyPage() {
    const createNewAccessToken = api.secrets.fetchSpotifyAccessToken.useMutation()

    return (
        <>
            <h1>Access Token data</h1>
            <button onClick={() => createNewAccessToken.mutate()}>Fetch new access Token</button>
            {createNewAccessToken.isLoading && <p>Loading...</p>}
            {createNewAccessToken.data && (
                <div>Access Token: <pre>{createNewAccessToken.data} </pre> </div>)}
        </>
    )
}