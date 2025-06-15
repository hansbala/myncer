import { useState } from "react"
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { Label } from "@/components/ui/label"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { useConnectedDatasources } from "@/hooks/useConnectedDatasources"

export const CreateSyncDialog = () => {
  const [open, setOpen] = useState(false)
  const { datasources: connectedDatasources, loading } = useConnectedDatasources()
  const [sourceDatasource, setSourceDatasource] = useState("")
  const [sourcePlaylist, setSourcePlaylist] = useState("")
  const [targetDatasource, setTargetDatasource] = useState("")
  const [targetPlaylist, setTargetPlaylist] = useState("")

  const handleCreateSync = () => {
    console.log("Syncing from:", {
      source: { datasource: sourceDatasource, playlist: sourcePlaylist },
      target: { datasource: targetDatasource, playlist: targetPlaylist },
    })
    setOpen(false)
  }

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Create Sync</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create a new sync</DialogTitle>
        </DialogHeader>

        <div className="flex flex-col space-y-6 py-2">
          {/* Source */}
          <div>
            <Label className="text-sm text-center">Source</Label>
            <div className="grid grid-cols-2 gap-4 mt-1">
              <Select onValueChange={setSourceDatasource}>
                <SelectTrigger>
                  <SelectValue placeholder="Datasource" />
                </SelectTrigger>
                <SelectContent>
                  {
                    connectedDatasources.map((ds) => (
                      <SelectItem key={ds} value={ds.toLowerCase()}>
                        {ds}
                      </SelectItem>
                    ))
                  }
                </SelectContent>
              </Select>

              <Select onValueChange={setSourcePlaylist}>
                <SelectTrigger>
                  <SelectValue placeholder="Playlist" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="liked">Liked Songs</SelectItem>
                  <SelectItem value="favorites">Favorites</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          {/* Target */}
          <div>
            <Label className="text-sm">Target</Label>
            <div className="grid grid-cols-2 gap-4 mt-1">
              <Select onValueChange={setTargetDatasource}>
                <SelectTrigger>
                  <SelectValue placeholder="Datasource" />
                </SelectTrigger>
                <SelectContent>
                  {
                    connectedDatasources.map((ds) => (
                      <SelectItem key={ds} value={ds.toLowerCase()}>
                        {ds}
                      </SelectItem>
                    ))
                  }
                </SelectContent>
              </Select>

              <Select onValueChange={setTargetPlaylist}>
                <SelectTrigger>
                  <SelectValue placeholder="Playlist" />
                </SelectTrigger>
                <SelectContent>
                  <SelectItem value="liked">Liked Songs</SelectItem>
                  <SelectItem value="favorites">Favorites</SelectItem>
                </SelectContent>
              </Select>
            </div>
          </div>

          <Button
            onClick={handleCreateSync}
            disabled={
              !sourceDatasource || !sourcePlaylist || !targetDatasource || !targetPlaylist
            }
            className="w-full"
          >
            Create Sync
          </Button>
        </div>
      </DialogContent>
    </Dialog>
  )
}

