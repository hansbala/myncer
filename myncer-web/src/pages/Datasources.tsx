import { PageWrapper } from "@/components/PageWrapper"
import { Button } from "@/components/ui/button"
import { getSpotifyAuthUrl } from "@/lib/utils"
import { ArrowRightIcon } from "lucide-react"

export const Datasources = () => {
  const handleConnectSpotify = () => {
    // window.location since we're redirecting externally.
    window.location.href = getSpotifyAuthUrl()
  }

  return (
  <PageWrapper>
    <div className="min-w-xl px-4 py-8 space-y-8">
      <div>
        <h1 className="text-3xl font-bold">Datasources</h1>
        <p className="text-muted-foreground mt-1 text-sm">
          Manage integrations with third-party services like Spotify.
        </p>
      </div>

      <div className="rounded-xl border bg-card p-6 shadow-sm">
        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-lg font-semibold">Spotify</h2>
            <p className="text-sm text-muted-foreground">
              Connect account to sync playlists.
            </p>
          </div>
          <Button onClick={handleConnectSpotify} className="flex items-center gap-1">
            Connect
            <ArrowRightIcon className="w-4 h-4" />
          </Button>
        </div>
      </div>
    </div>
  </PageWrapper>
)
}
