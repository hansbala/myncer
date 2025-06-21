import { createUser } from "@/generated_grpc/myncer/user-UserService_connectquery"
import { useMutation } from "@connectrpc/connect-query"
import { toast } from "sonner"

export const useCreateUser = () => {
  return useMutation(createUser, {
    onSuccess: () => {
      toast.success("User created!")
    },
    onError: (error) => {
      toast.error(`Failed to create user: ${error.rawMessage}`)
    }
  })
}
