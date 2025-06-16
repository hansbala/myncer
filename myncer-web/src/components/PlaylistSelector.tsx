import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import type { Playlist } from "@/generated_api/src"
import { Controller, type UseControllerProps, type FieldValues } from "react-hook-form"

interface PlaylistProps<T extends FieldValues> extends UseControllerProps<T> {
  playlists: Playlist[]
  label: string
  disabled?: boolean
}

export function PlaylistSelector<T extends FieldValues>({
  playlists,
  label,
  disabled,
  ...controllerProps
}: PlaylistProps<T>) {
  return (
    <Controller
      {...controllerProps}
      render={({ field }) => (
        <Select value={field.value} onValueChange={field.onChange} disabled={disabled}>
          <SelectTrigger>
            <SelectValue placeholder={label} />
          </SelectTrigger>
          <SelectContent>
            {playlists.map((p) => (
              <SelectItem key={p.playlistId} value={p.playlistId}>
                {p.name || p.playlistId}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      )}
    />
  )
}

