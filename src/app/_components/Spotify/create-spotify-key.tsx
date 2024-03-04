"use client";

import { useRouter } from "next/navigation";
import { useState } from "react";

import { api } from "~/trpc/react";
import Button from "../Button/Button";

export function CreateSpotifyKey() {
  const router = useRouter();
  const [modify, setModify] = useState<boolean>(false)
  const [clientId, setClientId] = useState("")
  const [clientSecret, setClientSecret] = useState("")
  const { data } = api.secrets.getSpotifySecret.useQuery()

  const createSpotifyApiKey = api.secrets.setSpotifySecret.useMutation({
    onSuccess: () => {
      router.refresh()
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
      <div className="flex flex-row items-center">
        <label className="flex-shrink-0 flex-grow-0 basis-1/6">Client ID:</label>
        <input
          type="text"
          id="client-id"
          placeholder="Client Id"
          value={modify ? clientId : data?.clientId}
          disabled={!modify}
          onChange={(e) => setClientId(e.target.value)}
          className="flex-grow flex-shrink-0 basis-5/6 rounded-md px-4 py-2 text-black"
        />
      </div>
      <div className="flex flex-row items-center">
        <label className="flex-shrink-0 flex-grow-0 basis-1/6">Client Secret:</label>
        <input
          type="text"
          placeholder="Client Secret"
          value={modify ? clientSecret : data?.clientSecret}
          disabled={!modify}
          onChange={(e) => setClientSecret(e.target.value)}
          className="flex-grow flex-shrink-0 basis-5/6 rounded-md px-4 py-2 text-black"
        />
      </div>
      <div className="flex flex-row gap-4">
        <Button type="button" onClick={(e) => setModify(!modify)}>Modify</Button>
        <Button type="submit" disabled={createSpotifyApiKey.isLoading}>
          {createSpotifyApiKey.isLoading ? "Saving..." : "Save"}
        </Button>
      </div>
    </form>
  );
}
