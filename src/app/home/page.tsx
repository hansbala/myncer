import { redirect } from 'next/navigation'
import { getServerAuthSession } from '~/server/auth'

async function HomePage() {
  const session = await getServerAuthSession()

  if (!session) {
    redirect('/')
  }

  return (
    <>
      <h1 className="text-5xl font-extrabold tracking-tight sm:text-[5rem]">
        Current Syncs:
      </h1>
    </>
  )
}

export default HomePage
