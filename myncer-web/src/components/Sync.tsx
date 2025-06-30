import { Button } from "@/components/ui/button"
import { Loader2, Trash2 } from "lucide-react"
import { OneWaySyncRender } from "./OneWaySyncRender"
import { useDeleteSync } from "@/hooks/useDeleteSync"
import { useRunSync } from "@/hooks/useRunSync"
import { SyncStatus, type Sync } from "@/generated_grpc/myncer/sync_pb"
import { protoTimestampToDate } from "@/lib/utils"
import { useListSyncRuns } from "@/hooks/useListSyncRuns"

export const SyncRender = ({ sync }: { sync: Sync }) => {
  const { runSync, isRunningSync } = useRunSync()
  const { deleteSync, isDeleting } = useDeleteSync()
  const { syncRuns } = useListSyncRuns()
  const isSyncRunning = syncRuns.some(
    (run) => run.syncId === sync.id && run.syncStatus === SyncStatus.RUNNING,
  )
  const { syncVariant, id, createdAt } = sync

  const renderVariantLabel = () => {
    switch (syncVariant.case) {
      case "oneWaySync":
        return "One-Way"
      default:
        return "Unknown"
    }
  }

  return (
    <div className="w-full rounded-lg border bg-muted p-4 shadow-sm space-y-4">
      {/* Header: Variant label + Actions */}
      <div className="flex items-center justify-between">
        <span className="text-xs text-muted-foreground">
          {renderVariantLabel()} Sync
        </span>
      </div>

      {/* Sync Details */}
      <div className="mb-8">
        {syncVariant.case === "oneWaySync" && (
          <OneWaySyncRender sync={syncVariant.value} />
        )}
      </div>

      {/* Footer: Metadata */}
      <div className="flex justify-between items-center">
        <div className="space-y-1 text-xs text-muted-foreground">
          <div>Last synced at coming soon…</div>
          <div>Created at {createdAt ? protoTimestampToDate(createdAt).toLocaleString() : "Unknown"}</div>
        </div>
        {/* Action Buttons */}
        <div className="flex items-center gap-2">
          {/* Delete Sync */}
          <Button
            size="sm"
            variant="destructive"
            onClick={() => deleteSync({ syncId: id })}
            disabled={isDeleting || isRunningSync || isSyncRunning}
          >
            {isDeleting ? (
              <>
                <Loader2 className="w-4 h-4 animate-spin mr-2" />
                Deleting
              </>
            ) : (
              <>
                <Trash2 className="w-4 h-4 mr-2" />
                Delete
              </>
            )}
          </Button>
          {/* Run Sync */}
          <Button
            size="sm"
            onClick={() => runSync({ syncId: id })}
            disabled={isRunningSync || isDeleting || isSyncRunning}
          >
            {(isRunningSync || isSyncRunning) ? (
              <>
                <Loader2 className="w-4 h-4 animate-spin mr-2" />
                Running
              </>
            ) : (
              "Run Sync"
            )}
          </Button>
        </div>
      </div>
    </div>
  )
}

