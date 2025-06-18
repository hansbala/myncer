import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserService } from "@/generated_grpc/myncer/user_pb";

const serverUrl = "http://localhost:6969"

export const useUserService = () => {
  const transport = createConnectTransport({
    baseUrl: serverUrl,
    fetch: (input: RequestInfo | URL, init?: RequestInit) => {
      return fetch(input, {
        ...init,
        credentials: "include",
      });
    },
  });

  const client = createClient(UserService, transport);
  return client;
};
