import { CreateSyncDialog } from "@/components/CreateSyncDialog"
import { PageWrapper } from "@/components/PageWrapper"
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
          <CreateSyncDialog />
        </div>

        <div className="rounded-xl border bg-card p-6 shadow-sm">
          <div className="space-y-4">
            {syncs.length === 0 ? (
              <p className="text-muted-foreground">No syncs found.</p>
            ) : (
              syncs.map((sync) => (
                <div key={sync.id} className="border-b py-4 last:border-0">
                  <h2 className="text-lg font-semibold">{sync.id}</h2>
                </div>
              ))
            )}
          </div>
        </div>
      </div>
    </PageWrapper>
  )
}
