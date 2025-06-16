import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Controller, type UseControllerProps, type FieldValues } from "react-hook-form"
import type { Datasource } from "@/generated_api/src"

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
          <SelectTrigger>
            <SelectValue placeholder={label} />
          </SelectTrigger>
          <SelectContent>
            {datasources.map((ds) => (
              <SelectItem key={ds} value={ds}>
                {ds}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      )}
    />
  )
}

