import type { Sync } from "@/generated_api/src"
import { useRunSync } from "@/hooks/useRunSync"
import { Button } from "@/components/ui/button"
import { Loader2 } from "lucide-react"
import { OneWaySyncRender } from "./OneWaySyncRender"

export const SyncRender = ({ sync }: { sync: Sync }) => {
  const { runSync, isRunningSync } = useRunSync()
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
      <div className="flex items-center justify-between">
        <span className="text-xs text-muted-foreground">{renderVariantLabel()} Sync</span>

        <Button
          size="sm"
          onClick={() => runSync(id)}
          disabled={isRunningSync}
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

      {syncData.syncVariant === "ONE_WAY" && (
        <OneWaySyncRender sync={syncData} />
      )}

      {syncData.syncVariant === "MERGE" && (
        <div className="text-sm text-muted-foreground italic">
          Merging playlists (details coming soon...)
        </div>
      )}

      <div className="space-y-1">
        <div className="text-xs text-muted-foreground">
          Last synced details coming soon...
        </div>
        <div className="text-xs text-muted-foreground">
          Created at {new Date(createdAt).toLocaleString()}
        </div>
      </div>
    </div>
  )
}
