import { useForm } from "react-hook-form"
import { z } from "zod"
import { zodResolver } from "@hookform/resolvers/zod"
import { useApiClient } from "../hooks/useApiClient"
import { Link, useNavigate } from "react-router-dom"
import { useAuth } from "../hooks/useAuth"

const signUpSchema = z
  .object({
    firstName: z
      .string()
      .min(3, "First name must be at least 3 characters")
      .max(10, "First name cannot be more than 10 characters"),
    lastName: z
      .string()
      .min(3, "Last name must be at least 3 characters")
      .max(10, "Last name cannot be more than 10 characters"),
    email: z.string().email({ message: "Enter a valid email" }),
    password: z.string().min(8, { message: "Password must be at least 8 characters" }),
    confirmPassword: z.string().min(8, { message: "Please confirm your password" }),
  })
  .refine((data) => data.password === data.confirmPassword, {
    message: "Passwords do not match",
    path: ["confirmPassword"],
  })

type SignUpFormInputs = z.infer<typeof signUpSchema>

export const SignUp = () => {
  const apiClient = useApiClient()
  const navigate = useNavigate()
  const { isAuthenticated } = useAuth()

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    setError,
  } = useForm<SignUpFormInputs>({
    resolver: zodResolver(signUpSchema),
  })

  const onSubmit = async (data: SignUpFormInputs) => {
    try {
      await apiClient.createUser({
        createUserRequest: {
          firstName: data.firstName,
          lastName: data.lastName,
          email: data.email,
          password: data.password,
        },
      })
      navigate("/")
    } catch (err) {
      setError("email", {
        type: "manual",
        message: "An account with this email may already exist",
      })
    }
  }

  if (isAuthenticated) {
    navigate("/")
  }

  return (
    <div className="flex h-screen w-screen items-center justify-center">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-full max-w-xl space-y-6 rounded-2xl border border-gray-200 p-10 shadow-md"
      >
        <h2 className="text-center text-3xl font-bold ">Create Account</h2>

        <div className="grid grid-cols-1 gap-4 sm:grid-cols-2">
          <div>
            <label className="mb-2 block text-base font-medium ">First Name</label>
            <input
              type="text"
              className="w-full rounded-md border px-4 py-3 shadow-sm focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
              {...register("firstName")}
            />
            {errors.firstName && (
              <p className="mt-1 text-sm text-red-600">{errors.firstName.message}</p>
            )}
          </div>

          <div>
            <label className="mb-2 block text-base font-medium ">Last Name</label>
            <input
              type="text"
              className="w-full rounded-md border px-4 py-3 shadow-sm focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
              {...register("lastName")}
            />
            {errors.lastName && (
              <p className="mt-1 text-sm text-red-600">{errors.lastName.message}</p>
            )}
          </div>
        </div>

        <div>
          <label className="mb-2 block text-base font-medium ">Email</label>
          <input
            type="email"
            className="w-full rounded-md border px-4 py-3 shadow-sm focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
            {...register("email")}
          />
          {errors.email && (
            <p className="mt-1 text-sm text-red-600">{errors.email.message}</p>
          )}
        </div>

        <div>
          <label className="mb-2 block text-base font-medium ">Password</label>
          <input
            type="password"
            className="w-full rounded-md border px-4 py-3 shadow-sm focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
            {...register("password")}
          />
          {errors.password && (
            <p className="mt-1 text-sm text-red-600">{errors.password.message}</p>
          )}
        </div>

        <div>
          <label className="mb-2 block text-base font-medium ">Confirm Password</label>
          <input
            type="password"
            className="w-full rounded-md border px-4 py-3 shadow-sm focus:border-primary focus:outline-none focus:ring-1 focus:ring-primary"
            {...register("confirmPassword")}
          />
          {errors.confirmPassword && (
            <p className="mt-1 text-sm text-red-600">{errors.confirmPassword.message}</p>
          )}
        </div>

        <button
          type="submit"
          disabled={isSubmitting}
          className="w-full rounded-md bg-primary px-6 py-3 text-lg text-secondary hover:bg-primary/90 disabled:cursor-not-allowed disabled:opacity-50"
        >
          {isSubmitting ? "Creating account..." : "Sign Up"}
        </button>

        <p className="text-center text-sm text-gray-600">
          Already have an account?{" "}
          <Link to="/login" className="text-primary hover:underline">
            Login here
          </Link>
        </p>
      </form>
    </div>
  )
}

