import { Loader2 } from "lucide-react"

export function Spinner({ className }: { className?: string }) {
  return <Loader2 className={`h-4 w-4 animate-spin text-muted-foreground ${className || ""}`} />
}
