import { Home, FolderSync, Settings } from "lucide-react"

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"
import { Separator } from "@radix-ui/react-separator"

// Menu items.
const items = [
  {
    title: "Dashboard",
    url: "#",
    icon: Home,
  },
  {
    title: "Syncs",
    url: "#",
    icon: FolderSync,
  },
  {
    title: "Settings",
    url: "#",
    icon: Settings,
  },
]

export const MyncerSidebar = () => {
  return (
    <Sidebar>
      <SidebarContent>
        <h1 className="w-full px-5 py-4 text-xl">Myncer</h1>
        <SidebarGroup>
          {/* <SidebarGroupLabel>Myncer</SidebarGroupLabel> */}
          <SidebarGroupContent>
            <SidebarMenu>
              {items.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton asChild>
                    <a href={item.url}>
                      <item.icon />
                      <span>{item.title}</span>
                    </a>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  )
}
