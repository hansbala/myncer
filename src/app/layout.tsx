import '~/styles/globals.css'

import { Inter } from 'next/font/google'
import { cookies } from 'next/headers'

import { TRPCReactProvider } from '~/trpc/react'
import Navbar from './_components/Navbar/Navbar'

const inter = Inter({
  subsets: ['latin'],
  variable: '--font-sans',
})

export const metadata = {
  title: 'Myncer',
  description: 'Music Syncer',
  icons: [{ rel: 'icon', url: '/favicon.ico' }],
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={`font-sans ${inter.variable}`}>
        <TRPCReactProvider cookies={cookies().toString()}>
          <Navbar />
          <main className="h-full w-full">{children}</main>
        </TRPCReactProvider>
      </body>
    </html>
  )
}
