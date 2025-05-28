import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "./hooks/useAuth";

export function RequireAuth({ children }: { children: React.ReactNode }) {
  const { isAuthenticated, loading } = useAuth();
  const location = useLocation();

  if (loading) {
    // Optional: Replace with a proper loading spinner
    return <div className="text-center p-4">Checking authentication...</div>;
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return <>{children}</>;
}
