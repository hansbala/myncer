import { Outlet } from "react-router-dom"
import { SidebarProvider } from "@/components/ui/sidebar"
import { MyncerSidebar } from "@/components/MyncerSidebar"
import { Toaster } from "sonner"

export const Root = () => {
  return (
    <SidebarProvider>
      <MyncerSidebar />
      <main className="w-full">
        <Outlet />
      </main>
      <Toaster />
    </SidebarProvider>
  )
}
