import { CreateSyncDialog } from "@/components/CreateSyncDialog"
import { PageWrapper } from "@/components/PageWrapper"

export const Syncs = () => {
  return (
    <PageWrapper>
      <div className="min-w-xl px-4 py-8 space-y-8">
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
          <p className="text-muted-foreground">
            Listing syncs is not yet implemented.
          </p>
        </div>
      </div>
    </PageWrapper>
  )
}
