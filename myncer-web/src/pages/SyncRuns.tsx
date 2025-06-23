import { PageWrapper } from "@/components/PageWrapper"
import { SyncRunRender } from "@/components/SyncRunRender"
import { PageLoader } from "@/components/ui/page-loader"
import { useListSyncRuns } from "@/hooks/useListSyncRuns"

export const SyncRuns = () => {
  const { syncRuns, loading } = useListSyncRuns()
  if (loading) {
    return <PageLoader />
  }
  return (
    <PageWrapper>
      <div className="max-w-5xl px-4 py-8 space-y-8">
        <div className="flex items-center justify-between">
          <div>
            <h1 className="text-3xl font-bold">Syncs Runs</h1>
            <p className="text-muted-foreground mt-1 text-sm">
              All sync runs are listed here.
            </p>
          </div>
        </div>

        <div className="space-y-4">
          {loading && (
              <div className="rounded-xl border bg-card p-6 shadow-sm">
                <p className="text-muted-foreground">Loading sync runs...</p>
              </div>
            )
          }
          {!loading && syncRuns.length === 0 ? (
            <div className="rounded-xl border bg-card p-6 shadow-sm">
              <p className="text-muted-foreground">Run a sync to see it here.</p>
            </div>
          ) : (
            syncRuns.map((syncRun) => <SyncRunRender key={syncRun.runId} syncRun={syncRun} />)
          )}
        </div>
      </div>
    </PageWrapper>
  )
}
