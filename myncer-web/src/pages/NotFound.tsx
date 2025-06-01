import { Button } from "@/components/ui/button"
import { useNavigate } from "react-router-dom"

export const NotFound = () => {
  const navigate = useNavigate()

  return (
    <div className="w-screen h-screen flex flex-col justify-center items-center space-y-4">
      <h1 className="text-lg">Page not found</h1>
      <Button onClick={() => navigate("/")}>Back to home</Button>
    </div>
  )
}
