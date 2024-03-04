import { redirect } from "next/navigation";

import { CreateSpotifyKey } from "~/app/_components/Spotify/create-spotify-key";
import { getServerAuthSession } from "~/server/auth";
import { api } from "~/trpc/server";

export default async function Home() {
  const session = await getServerAuthSession();

  if (session) {
    redirect('/home')
  }

  return (
    <main className="flex min-h-screen flex-col items-center justify-center">
      <div className="container flex flex-col items-center justify-center gap-12 px-4 py-16 ">
        <h1 className="text-5xl font-extrabold tracking-tight sm:text-[5rem]">
          Myncer
        </h1>
        <div>🔄 The OSS music syncer 🔄</div>

        <CrudShowcase />
      </div>
    </main>
  );
}

async function CrudShowcase() {
  const session = await getServerAuthSession();
  if (!session?.user) return null;

  const spotifyApiKey = await api.secrets.getSpotifySecret.query();

  return (
    <div className="w-full max-w-xs">
      {spotifyApiKey ? (
        <p className="truncate">Client ID: {spotifyApiKey.clientId} | Client Secret: {spotifyApiKey.clientSecret} </p>
      ) : (
        <p>You have no posts yet.</p>
      )}

      <CreateSpotifyKey />
    </div>
  );
}
