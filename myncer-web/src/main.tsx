import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { TransportProvider } from "@connectrpc/connect-query"

import App from './App.tsx'
import './index.css'
import { ThemeProvider } from './components/ThemeProvider.tsx'
import { createConnectTransport } from '@connectrpc/connect-web'

const GRPC_SERVER_URL = import.meta.env.VITE_GRPC_BASE_URL || 'https://myncer-api.hansbala.com'

const queryClient = new QueryClient()
const connectTransport = createConnectTransport({
  baseUrl: GRPC_SERVER_URL,
  // For cookie auth.
  fetch: (input: RequestInfo | URL, init?: RequestInit) => {
    return fetch(input, {
      ...init,
      credentials: "include",
    });
  },
})

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider>
      <TransportProvider transport={connectTransport}>
        <QueryClientProvider client={queryClient}>
          <App />
        </QueryClientProvider>
      </TransportProvider>
    </ThemeProvider>
  </StrictMode>,
)
