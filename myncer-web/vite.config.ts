import path from "path"
import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks(id) {
          if (id.includes("node_modules")) {
            if (id.includes("react")) return "react"
            if (id.includes("@radix-ui")) return "radix"
            if (id.includes("lucide-react")) return "icons"
            if (id.includes("connect-query") || id.includes("protobuf")) return "connectrpc"
            if (id.includes("react-query")) return "query"
            if (id.includes("zod") || id.includes("react-hook-form")) return "forms"
          }
        },
      },
    },
  }
})
