import Link from 'next/link'
import { RxHamburgerMenu } from 'react-icons/rx'
import { getServerAuthSession } from '~/server/auth'

export default async function Navbar() {
  const session = await getServerAuthSession()

  return (
    <div className="navbar sticky top-0 border-b bg-white text-black">
      <div className="navbar-start">
        <div className="dropdown">
          <div tabIndex={0} role="button" className="btn btn-ghost lg:hidden">
            <RxHamburgerMenu className="h-5 w-5" />
          </div>
          <ul
            tabIndex={0}
            className="menu dropdown-content menu-sm z-[1] mt-3 w-52 rounded-box bg-white p-2 text-black shadow"
          >
            <li>
              <Link href="/home">Home</Link>
            </li>
            <li>
              <Link href="/secrets">Secrets</Link>
            </li>
            <li>
              <Link href="/spotify">Spotify</Link>
            </li>
            <li>
              <Link href="/youtube">Youtube</Link>
            </li>
            <li>
              <Link href="/newsync">New Sync</Link>
            </li>
          </ul>
        </div>
        <a className="btn btn-ghost text-xl">Myncer</a>
      </div>
      <div className="navbar-center hidden lg:flex">
        <ul className="menu menu-horizontal px-1">
          <li>
            <Link href="/home">Home</Link>
          </li>
          <li>
            <Link href="/secrets">Secrets</Link>
          </li>
          <li>
            <Link href="/spotify">Spotify</Link>
          </li>
          <li>
            <Link href="/youtube">Youtube</Link>
          </li>
          <li>
            <Link href="/newsync">New Sync</Link>
          </li>
        </ul>
      </div>
      <div className="navbar-end">
        {session ? (
          <Link href="/api/auth/signout" className="btn">
            Signout
          </Link>
        ) : (
          <Link href="/api/auth/signin" className="btn">
            Sign In
          </Link>
        )}
      </div>
    </div>
  )
}
