import { redirect } from 'next/navigation'
import { getServerAuthSession } from '~/server/auth'
import { api } from '~/trpc/server'

async function HomePage() {
  const session = await getServerAuthSession()
  const syncs = await api.syncs.getSyncs.query()

  if (!session) {
    redirect('/')
  }

  return (
    <>
      <h1 className="text-5xl font-extrabold tracking-tight sm:text-[5rem]">
        Current Syncs:
      </h1>
      <div>
        {syncs.map((sync) => {
          return (
            <div key={sync.id}>
              <p>{sync.sourcePlaylistId}</p>
              <p>{sync.destinationPlaylistId}</p>
            </div>
          )
        })}
      </div>
    </>
  )
}

export default HomePage
