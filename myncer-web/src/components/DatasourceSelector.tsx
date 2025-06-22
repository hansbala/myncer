import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Datasource } from "@/generated_grpc/myncer/datasource_pb"
import { getDatasourceLabel } from "@/lib/utils"
import { Controller, type UseControllerProps, type FieldValues } from "react-hook-form"
import { toast } from "sonner"

interface Props<T extends FieldValues> extends UseControllerProps<T> {
  datasources: Datasource[]
  label: string
}

export function DatasourceSelector<T extends FieldValues>({
  datasources,
  label,
  ...controllerProps
}: Props<T>) {
  return (
    <Controller
      {...controllerProps}
      render={({ field }) => (
        <Select
          value={field.value != null ? String(field.value) : ""}
          onValueChange={(val) => {
            const numericVal = Number(val)
            if (!Object.values(Datasource).includes(numericVal)) {
              toast.error(`Unsupported datasource: ${val}`)
              return
            }
            field.onChange(numericVal)
          }}
        >
          <SelectTrigger className="w-full max-w-full">
            <SelectValue placeholder={label} />
          </SelectTrigger>
          <SelectContent>
            {datasources.map((ds) => {
              const datasourceLabel = getDatasourceLabel(ds)
              return (
                <SelectItem key={ds} value={String(ds)}>
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
