import { useState } from "react"
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Button } from "@/components/ui/button"
import { useForm } from "react-hook-form"
import { useConnectedDatasources } from "@/hooks/useConnectedDatasources"
import { usePlaylists } from "@/hooks/usePlaylists"
import type { Datasource } from "@/generated_api/src"
import { DatasourceSelector } from "./DatasourceSelector"
import { PlaylistSelector } from "./PlaylistSelector"
import { useCreateSync } from "@/hooks/useCreateSync"
import { Loader2 } from "lucide-react"

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
    formState: { isValid },
  } = useForm<FormValues>({
    mode: "onChange",
  })
  const { mutate: createSync, isPending: creating } = useCreateSync()

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
    createSync({
      syncVariant: "ONE_WAY",
      source: {
        datasource: data.sourceDatasource,
        playlistId: data.sourcePlaylistId,
      },
      destination: {
        datasource: data.targetDatasource,
        playlistId: data.targetPlaylistId,
      },
      // TODO: Add overwrite existing? to form and then use here.
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
      <DialogContent aria-describedby="create a one way sync">
        <DialogHeader>
          <DialogTitle>Create One-Way sync</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit(onSubmit)} className="space-y-6 py-2">
          <div className="flex flex-col space-y-2">
            <div className="grid grid-cols-[1fr_auto_1fr] items-center gap-4">
              {/* Source */}
              <div className="w-[200px] flex flex-col space-y-2">
                <DatasourceSelector<FormValues>
                  name="sourceDatasource"
                  control={control}
                  datasources={connectedDatasources}
                  label="Source Datasource"
                />
                <PlaylistSelector<FormValues>
                  name="sourcePlaylistId"
                  control={control}
                  playlists={sourcePlaylists}
                  label="Source Playlist"
                  disabled={!sourceDatasource}
                />
              </div>

              {/* Arrow */}
              <div className="text-center text-2xl text-muted-foreground">
                →
              </div>

              {/* Target */}
              <div className="w-[200px] flex flex-col space-y-2">
                <DatasourceSelector<FormValues>
                  name="targetDatasource"
                  control={control}
                  datasources={connectedDatasources}
                  label="Target Datasource"
                />
                <PlaylistSelector<FormValues>
                  name="targetPlaylistId"
                  control={control}
                  playlists={targetPlaylists}
                  label="Target Playlist"
                  disabled={!targetDatasource}
                />
              </div>
            </div>
          </div>

          <Button
            type="submit"
            disabled={!isValid || isFormLoading || creating}
            className="w-full"
          >
            {(isFormLoading || creating) ? (
              <div className="flex items-center justify-center space-x-2">
                <Loader2 className="h-4 w-4 animate-spin" />
                <span>{isFormLoading ? "Loading..." : "Creating..."}</span>
              </div>
            ) : (
              "Create Sync"
            )}
          </Button>
        </form>
      </DialogContent>
    </Dialog>
  )
}

