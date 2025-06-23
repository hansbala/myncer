import { User, LogOut, Database, FolderSync, Play } from "lucide-react"
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import { useAuth } from "@/hooks/useAuth"

const menuItems = [
  {
    title: "Syncs",
    url: "/syncs",
    icon: FolderSync,
  },
  {
    title: "Sync Runs",
    url: "/syncruns",
    icon: Play,
  },
  {
    title: "Datasources",
    url: "/datasources",
    icon: Database,
  },
]

export const MyncerSidebar = () => {
  const { user, loading, isAuthenticated, logout } = useAuth()

  return (
    <Sidebar>
      <SidebarContent className="flex h-full flex-col justify-between">
        {/* Top section */}
        <div>
          <h1 className="w-full px-5 py-4 text-xl font-semibold">Myncer</h1>
          <SidebarGroup>
            <SidebarGroupContent>
              <SidebarMenu>
                {menuItems.map((item) => (
                  <SidebarMenuItem key={item.title}>
                    <SidebarMenuButton asChild>
                      <a href={item.url} className="flex items-center gap-2">
                        <item.icon className="w-4 h-4" />
                        <span>{item.title}</span>
                      </a>
                    </SidebarMenuButton>
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        </div>

        {/* Dark mode toggle */}
        {/*
        <div className="flex items-center justify-center border-t px-4 py-3 text-sm">
          <ThemeToggle />
        </div>
        */}

        {/* User section */}
        {isAuthenticated && !loading && user && (
          <div className="flex items-center justify-between border-t px-4 py-3 text-sm">
            <div className="flex items-center gap-2">
              <User className="h-4 w-4" />
              <span>{user.firstName || user.email}</span>
            </div>
            <button
              onClick={() => logout.mutateAsync({})}
              className="flex items-center gap-2 text-red-500 hover:underline"
            >
              <LogOut className="h-4 w-4" />
              <span>Logout</span>
            </button>
          </div>
        )}
      </SidebarContent>
    </Sidebar>
  )
}

