import { redirect } from "next/navigation"
import { getServerAuthSession } from "~/server/auth"
import AuthorizeSpotify from "../_components/Spotify/AuthorizeSpotify"
import AuthorizeGoogle from "../_components/Google/AuthorizeGoogle"

export default async function SecretsPage() {
  const session = await getServerAuthSession()
  if (!session) {
    redirect('/')
  }

  return (
    <div className="p-8">
      <h1 className="text-center mb-10 text-xl">Authenticate with apps below to start syncs</h1>
      <div className="flex flex-col gap-10">
        <div className="w-1/2 border rounded-lg p-3">
          <AuthorizeSpotify />
        </div>
        <div className="w-1/2 border rounded-lg p-3">
          <AuthorizeGoogle />
        </div>
      </div>
    </div>
  )
}