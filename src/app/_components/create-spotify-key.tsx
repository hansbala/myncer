"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";

import { api } from "~/trpc/react";

export function CreateSpotifyKey() {
  const router = useRouter();
  const [clientId, setClientId] = useState("")
  const [clientSecret, setClientSecret] = useState("")

  const createSpotifyApiKey = api.secrets.setSpotifySecret.useMutation({
    onSuccess: () => {
      router.refresh()
      setClientId("")
      setClientSecret("")
    }
  })

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        createSpotifyApiKey.mutate({ clientId, clientSecret });
      }}
      className="flex flex-col gap-2"
    >
      <input
        type="text"
        placeholder="Client Id"
        value={clientId}
        onChange={(e) => setClientId(e.target.value)}
        className="w-full rounded-full px-4 py-2 text-black"
      />
      <input
        type="text"
        placeholder="Client Secret"
        value={clientSecret}
        onChange={(e) => setClientSecret(e.target.value)}
        className="w-full rounded-full px-4 py-2 text-black"
      />
      <button
        type="submit"
        className="rounded-full bg-white/10 px-10 py-3 font-semibold transition hover:bg-white/20"
        disabled={createSpotifyApiKey.isLoading}
      >
        {createSpotifyApiKey.isLoading ? "Submitting..." : "Submit"}
      </button>
    </form>
  );
}
