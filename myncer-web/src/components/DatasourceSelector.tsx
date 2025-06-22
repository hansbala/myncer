import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import { Controller, type UseControllerProps, type FieldValues } from "react-hook-form"

interface Props<T extends FieldValues> extends UseControllerProps<T> {
  datasources: Datasource[]
  label: string
}

export function DatasourceSelector<T extends FieldValues>({ datasources, label, ...controllerProps }: Props<T>) {
  return (
    <Controller
      {...controllerProps}
      render={({ field }) => (
        <Select value={field.value} onValueChange={field.onChange}>
          <SelectTrigger className="w-full max-w-full">
            <SelectValue placeholder={label} />
          </SelectTrigger>
          <SelectContent>
            {datasources.map((ds) => {
              let datasourceLabel: string
              switch (ds) {
                case Datasource.SPOTIFY:
                  datasourceLabel = "Spotify"
                  break
                case Datasource.YOUTUBE:
                  datasourceLabel = "YouTube"
                  break
                default:
                  datasourceLabel = ds.toString()
              }
              return (
                <SelectItem key={ds} value={datasourceLabel}>
                  {datasourceLabel}
                </SelectItem>
              )
            })}
          </SelectContent>
        </Select>
      )}
    />
  )
}

