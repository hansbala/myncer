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
import { Controller, useForm } from "react-hook-form"
import { useConnectedDatasources } from "@/hooks/useConnectedDatasources"
import { usePlaylists } from "@/hooks/usePlaylists"
import type { Datasource } from "@/generated_api/src"
import { DatasourceSelector } from "./DatasourceSelector"
import { PlaylistSelector } from "./PlaylistSelector"

type FormValues = {
  sourceDatasource: Datasource
  sourcePlaylistId: string
  targetDatasource: Datasource
  targetPlaylistId: string
}

export const CreateOneWaySyncDialog = () => {
  const [open, setOpen] = useState(false)
  const { datasources: connectedDatasources, loading: datasourcesLoading } = useConnectedDatasources()

  const {
    control,
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
        <Button>Create One-Way Sync</Button>
      </DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create a new sync</DialogTitle>
        </DialogHeader>

        <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col space-y-6 py-2">
          {/* Source */}
          <div>
            <div className="grid grid-cols-2 gap-4 mt-1">
              <DatasourceSelector
                name="sourceDatasource"
                control={control}
                datasources={connectedDatasources}
                label="Source Datasource"
              />
              <PlaylistSelector
                name="sourcePlaylistId"
                control={control}
                playlists={sourcePlaylists}
                label="Source Playlist"
                disabled={!sourceDatasource}
              />
            </div>
          </div>

          {/* Target */}
          <div>
            <div className="grid grid-cols-2 gap-4 mt-1">
              <DatasourceSelector
                name="targetDatasource"
                control={control}
                datasources={connectedDatasources}
                label="Target Datasource"
              />
              <PlaylistSelector
                name="targetPlaylistId"
                control={control}
                playlists={targetPlaylists}
                label="Target Playlist"
                disabled={!targetDatasource}
              />
              <Select onValueChange={(val) => setValue("targetPlaylistId", val)} disabled={!targetDatasource}>
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

