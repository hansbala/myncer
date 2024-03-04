import Link from "next/link"
import { getServerAuthSession } from "~/server/auth"

export default async function Navbar() {
  const session = await getServerAuthSession()

  return (
    <nav className="flex flex-row justify-between p-5 bg-white text-black">
      <div className="flex flex-row gap-5">
        <Link href="/home" className="underline">
          Home
        </Link>
        <Link href="/secrets" className="underline">
          Secrets
        </Link>
      </div>
      {!session && (
        <Link href="/api/auth/signin" className="underline">Sign In</Link>
      )}
      {session && (
        <Link href="/api/auth/signout" className="underline">Signout</Link>
      )}
    </nav>
  )

}