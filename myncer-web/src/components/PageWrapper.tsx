import type React from "react"

export const PageWrapper = ({ children }: { children: React.ReactNode }) => {
  return <div className="p-5">{children}</div>
}
