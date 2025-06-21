import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { UserService } from "@/generated_grpc/myncer/user_pb";
import { GRPC_SERVER_URL } from "./constants";

export const useUserService = () => {
  const transport = createConnectTransport({
    baseUrl: GRPC_SERVER_URL,
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
