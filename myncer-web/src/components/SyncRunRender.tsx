import { SyncStatus, type SyncRun } from "@/generated_grpc/myncer/sync_pb"
import { protoTimestampToDate } from "@/lib/utils"
import { Button } from "./ui/button"
import { Link } from "react-router-dom"
import clsx from "clsx"

const syncStatusToUI: Record<
  SyncStatus,
  { label: string; className: string }
> = {
  [SyncStatus.UNSPECIFIED]: {
    label: "Unspecified",
    className: "bg-gray-200 text-gray-700",
  },
  [SyncStatus.PENDING]: {
    label: "Pending",
    className: "bg-yellow-100 text-yellow-800",
  },
  [SyncStatus.RUNNING]: {
    label: "Running",
    className: "bg-blue-100 text-blue-800 animate-pulse",
  },
  [SyncStatus.COMPLETED]: {
    label: "Completed",
    className: "bg-green-100 text-green-800",
  },
  [SyncStatus.FAILED]: {
    label: "Failed",
    className: "bg-red-100 text-red-800",
  },
  [SyncStatus.CANCELLED]: {
    label: "Cancelled",
    className: "bg-muted text-muted-foreground",
  },
}

interface SyncRunRenderProps {
  syncRun: SyncRun
}
export const SyncRunRender = ({ syncRun }: SyncRunRenderProps) => {
  const createdAt = syncRun.createdAt
    ? protoTimestampToDate(syncRun.createdAt).toLocaleString()
    : "Unknown"
  const updatedAt = syncRun.updatedAt
    ? protoTimestampToDate(syncRun.updatedAt).toLocaleString()
    : "Unknown"

  const status = syncStatusToUI[syncRun.syncStatus] ?? syncStatusToUI[SyncStatus.UNSPECIFIED]

  return (
    <div className="rounded-xl border bg-card p-6 shadow-sm">
      <div className="flex items-center justify-between">
        <div>
          <h2 className="text-lg font-semibold">Run ID: {syncRun.runId}</h2>
          <div
            className={clsx(
              "inline-block rounded-full px-3 py-0.5 text-xs font-medium mt-2",
              status.className
            )}
          >
            {status.label}
          </div>
          <p className="text-sm text-muted-foreground mt-2">Created: {createdAt}</p>
          <p className="text-sm text-muted-foreground">Updated: {updatedAt}</p>
        </div>
        <Button variant="outline" asChild>
          <Link to={`/syncs/${syncRun.syncId}`}>View Sync</Link>
        </Button>
      </div>
    </div>
  )
}
