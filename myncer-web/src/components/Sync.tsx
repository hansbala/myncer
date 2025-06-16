import type { Sync } from "@/generated_api/src"
import { useRunSync } from "@/hooks/useRunSync"
import { Button } from "@/components/ui/button"
import { Loader2 } from "lucide-react"

export const SyncRender = ({ sync }: { sync: Sync }) => {
  const { runSync, isRunningSync } = useRunSync()
  const { syncData, id, createdAt } = sync

  return (
    <div className="rounded-lg border bg-muted p-4 shadow-sm space-y-2">
      <div className="flex items-center justify-between items-center">
        <div className="space-y-1">
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium text-muted-foreground">Sync ID</span>
            <span className="text-xs text-muted-foreground ml-2">{id.slice(0, 8)}...</span>
          </div>
          <div className="text-xs text-muted-foreground italic">
            {syncData.syncVariant === "ONE_WAY"
              ? "One-Way Sync"
              : syncData.syncVariant === "MERGE"
                ? "Merge Sync"
                : "Unknown Variant"}
          </div>
        </div>

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
        <>
          <div className="text-sm">
            <span className="font-semibold">{syncData.source.datasource}</span> →{" "}
            <span className="font-semibold">{syncData.destination.datasource}</span>
          </div>
          {syncData.overwriteExisting && (
            <div className="text-xs text-yellow-800 bg-yellow-100 inline-block px-2 py-0.5 rounded">
              Overwrites destination
            </div>
          )}
        </>
      )}

      {syncData.syncVariant === "MERGE" && (
        <div className="text-sm text-muted-foreground italic">
          Merging playlists (details coming soon...)
        </div>
      )}

      <div className="text-xs text-muted-foreground">
        Created: {new Date(createdAt).toLocaleString()}
      </div>
    </div>
  )
}

