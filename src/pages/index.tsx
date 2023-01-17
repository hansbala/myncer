import Head from 'next/head'
import Image from 'next/image'

export default function Home() {
  return (
    <>
      <Head>
        <title>Myncer</title>
        <meta name="description" content="OSS music transfer and syncer utility" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <main>
        <h1 className="text-3xl font-bold underline">
          Myncer
        </h1>
      </main>
    </>
  )
}
