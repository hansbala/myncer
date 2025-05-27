import { Outlet } from "react-router-dom"

export const Root = () => {
  return (
    <div>
      <header style={{ padding: "1rem", borderBottom: "1px solid #ccc" }}>
        <h1>Myncer Syncing App</h1>
      </header>
      <main>
        <Outlet />
      </main>
    </div>
  )
}
