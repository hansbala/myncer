import { redirect } from 'next/navigation'
import { getServerAuthSession } from '~/server/auth'
import AuthorizeSpotify from '../_components/Spotify/AuthorizeSpotify'
import AuthorizeGoogle from '../_components/Google/AuthorizeGoogle'

export default async function SecretsPage() {
  const session = await getServerAuthSession()
  if (!session) {
    redirect('/')
  }

  return (
    <div className="container p-8">
      <h1 className="mb-10 text-xl">
        Authenticate with apps below to start syncs
      </h1>
      <div className="grid grid-cols-3 gap-4">
        <AuthorizeSpotify />
        <AuthorizeGoogle />
      </div>
    </div>
  )
}
