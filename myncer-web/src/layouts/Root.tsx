import { Outlet } from "react-router-dom"

export const Root = () => {
  return (
    <div>
      <header style={{ padding: "1rem", borderBottom: "1px solid #ccc" }}>
        <h1>🏠 Myncer App</h1>
      </header>

      <main style={{ padding: "2rem" }}>
        <Outlet />
      </main>

      <footer style={{ padding: "1rem", borderTop: "1px solid #ccc" }}>
        <small>© {new Date().getFullYear()} Myncer Inc.</small>
      </footer>
    </div>
  )
}
