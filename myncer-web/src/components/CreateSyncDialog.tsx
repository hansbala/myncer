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
import { useForm } from "react-hook-form"
import { useConnectedDatasources } from "@/hooks/useConnectedDatasources"
import { usePlaylists } from "@/hooks/usePlaylists"
import type { Datasource } from "@/generated_api/src"

type FormValues = {
  sourceDatasource: Datasource
  sourcePlaylistId: string
  targetDatasource: Datasource
  targetPlaylistId: string
}

export const CreateSyncDialog = () => {
  const [open, setOpen] = useState(false)
  const { datasources: connectedDatasources, loading: datasourcesLoading } = useConnectedDatasources()

  const {
    watch,
    handleSubmit,
    setValue,
    formState: { isValid },
  } = useForm<FormValues>({
    mode: "onChange",
  })

  const sourceDatasource = watch("sourceDatasource")
  const targetDatasource = watch("targetDatasource")

  const {
    playlists: sourcePlaylists,
    loading: sourcePlaylistsLoading,
  } = usePlaylists({ datasource: sourceDatasource })

  const {
    playlists: targetPlaylists,
    loading: targetPlaylistsLoading,
  } = usePlaylists({ datasource: targetDatasource })

  const onSubmit = (data: FormValues) => {
    console.log("Syncing from:", {
      source: {
        datasource: data.sourceDatasource,
        playlistId: data.sourcePlaylistId,
      },
      target: {
        datasource: data.targetDatasource,
        playlistId: data.targetPlaylistId,
      },
    })
    setOpen(false)
  }

  const isFormLoading =
    datasourcesLoading || sourcePlaylistsLoading || targetPlaylistsLoading

  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogTrigger asChild>
        <Button>Create Sync</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create a new sync</DialogTitle>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col space-y-6 py-2">
          {/* Source */}
          <div>
            <Label className="text-sm text-center">Source</Label>
            <div className="grid grid-cols-2 gap-4 mt-1">
              <Select onValueChange={(val) => setValue("sourceDatasource", val as Datasource)}>
                <SelectTrigger>
                  <SelectValue placeholder="Datasource" />
                </SelectTrigger>
                <SelectContent>
                  {connectedDatasources.map((ds) => (
                    <SelectItem key={ds} value={ds}>
                      {ds}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select onValueChange={(val) => setValue("sourcePlaylistId", val)}>
                <SelectTrigger>
                  <SelectValue placeholder="Playlist" />
                </SelectTrigger>
                <SelectContent>
                  {sourcePlaylists.map((p) => (
                    <SelectItem key={p.playlistId} value={p.playlistId}>
                      {p.name || p.playlistId}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          {/* Target */}
          <div>
            <Label className="text-sm text-center">Target</Label>
            <div className="grid grid-cols-2 gap-4 mt-1">
              <Select onValueChange={(val) => setValue("targetDatasource", val as Datasource)}>
                <SelectTrigger>
                  <SelectValue placeholder="Datasource" />
                </SelectTrigger>
                <SelectContent>
                  {connectedDatasources.map((ds) => (
                    <SelectItem key={ds} value={ds}>
                      {ds}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>

              <Select onValueChange={(val) => setValue("targetPlaylistId", val)}>
                <SelectTrigger>
                  <SelectValue placeholder="Playlist" />
                </SelectTrigger>
                <SelectContent>
                  {targetPlaylists.map((p) => (
                    <SelectItem key={p.playlistId} value={p.playlistId || ""}>
                      {p.name || p.playlistId}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
            </div>
          </div>

          <Button
            type="submit"
            disabled={!isValid || isFormLoading}
            className="w-full"
          >
            {isFormLoading ? "Loading..." : "Create Sync"}
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  )
}

