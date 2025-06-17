import type { Sync } from "@/generated_api/src"
import { useRunSync } from "@/hooks/useRunSync"
import { useDeleteSync } from "@/hooks/useDeleteSync"
import { Button } from "@/components/ui/button"
import { Loader2, Trash2 } from "lucide-react"
import { OneWaySyncRender } from "./OneWaySyncRender"

export const SyncRender = ({ sync }: { sync: Sync }) => {
  const { runSync, isRunningSync } = useRunSync()
  const { deleteSync, isDeleting } = useDeleteSync()
  const { syncData, id, createdAt } = sync

  const renderVariantLabel = () => {
    return syncData.syncVariant === "ONE_WAY"
      ? "One-Way"
      : syncData.syncVariant === "MERGE"
        ? "Merge"
        : "Unknown"
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
        {syncData.syncVariant === "ONE_WAY" && (
          <OneWaySyncRender sync={syncData} />
        )}
        {syncData.syncVariant === "MERGE" && (
          <div className="text-sm text-muted-foreground italic">
            Merging playlists (details coming soon…)
          </div>
        )}
      </div>

      {/* Footer: Metadata */}
      <div className="flex justify-between items-center">
        <div className="space-y-1 text-xs text-muted-foreground">
          <div>Last synced at coming soon…</div>
          <div>Created at {new Date(createdAt).toLocaleString()}</div>
        </div>
        {/* Action Buttons */}
        <div className="flex items-center gap-2">
          {/* Delete Sync */}
          <Button
            size="sm"
            variant="destructive"
            onClick={() => deleteSync(id)}
            disabled={isDeleting || isRunningSync}
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
            onClick={() => runSync(id)}
            disabled={isRunningSync || isDeleting}
          >
            {isRunningSync ? (
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

