import { useEffect } from "react"

export function ThemeProvider({ children }: { children: React.ReactNode }) {
  useEffect(() => {
    const media = window.matchMedia("(prefers-color-scheme: dark)")

    const apply = () => {
      const prefersDark = media.matches
      document.documentElement.classList.toggle("dark", prefersDark)
    }

    apply()
    media.addEventListener("change", apply)
    return () => media.removeEventListener("change", apply)
  }, [])

  return <>{children}</>
}
