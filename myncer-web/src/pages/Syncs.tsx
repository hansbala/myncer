import { CreateOneWaySyncDialog } from "@/components/CreateOneWaySyncDialog"
import { PageWrapper } from "@/components/PageWrapper"
import { SyncRender } from "@/components/Sync"
import { PageLoader } from "@/components/ui/page-loader"
import { useSyncs } from "@/hooks/useSyncs"

export const Syncs = () => {
  const { syncs, loading } = useSyncs()
  if (loading) {
    return <PageLoader />
  }
  return (
    <PageWrapper>
      <div className="max-w-5xl px-4 py-8 space-y-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Syncs</h1>
            <p className="text-muted-foreground mt-1 text-sm">
              Manage your data synchronization settings.
            </p>
          </div>
          <CreateOneWaySyncDialog />
        </div>

        <div className="space-y-4">
          {syncs.length === 0 ? (
            <div className="rounded-xl border bg-card p-6 shadow-sm">
              <p className="text-muted-foreground">No syncs found.</p>
            </div>
          ) : (
            syncs.map((sync) => <SyncRender sync={sync} />)
          )}
        </div>
      </div>
    </PageWrapper>
  )
}
