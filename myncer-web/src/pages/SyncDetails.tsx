import { PageWrapper } from "@/components/PageWrapper"
import { SyncRender } from "@/components/Sync"
import { PageLoader } from "@/components/ui/page-loader"
import { useSync } from "@/hooks/useSync"
import { useParams } from "react-router-dom"

export const SyncDetails = () => {
  const { syncId } = useParams()
  const { sync, loading } = useSync(syncId)
  if (!syncId || !sync) {
    return <div className="text-red-500">Sync ID is required</div>
  }
  if (loading) {
    return <PageLoader />
  }
  return (
    <PageWrapper>
      <div className="max-w-5xl px-4 py-8 space-y-8">
        <SyncRender key={sync.id} sync={sync} />
      </div>
    </PageWrapper>
  )
}
