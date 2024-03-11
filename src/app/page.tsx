import { redirect } from 'next/navigation'

import { getServerAuthSession } from '~/server/auth'

export default async function Home() {
  const session = await getServerAuthSession()

  if (session) {
    redirect('/home')
  }

  return (
    <div className="flex h-full w-full flex-col items-center justify-center gap-12 overflow-y-hidden">
      <h1 className="text-5xl font-extrabold tracking-tight sm:text-[5rem]">
        Myncer
      </h1>
      <div>🔄 The OSS music syncer 🔄</div>
    </div>
  )
}
