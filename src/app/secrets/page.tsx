import { redirect } from "next/navigation"
import { getServerAuthSession } from "~/server/auth"
import AuthorizeApp from "../_components/Spotify/AuthorizeApp"

export default async function SecretsPage() {
  const session = await getServerAuthSession()
  if (!session) {
    redirect('/')
  }

  return (
    <div className="p-8">
      <div className="text-center mb-10">Enter all secrets below to start syncing apps</div>
      <div className="w-1/2 border rounded-lg p-3">
        <h2 className="text-center mb-5">Spotify New Flow</h2>
        <AuthorizeApp />
      </div>
    </div>
  )
}