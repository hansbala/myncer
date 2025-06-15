import { Navigate, useLocation } from "react-router-dom";
import { useAuth } from "./hooks/useAuth";
import { PageLoader } from "./components/ui/page-loader";

export function RequireAuth({ children }: { children: React.ReactNode }) {
  const { isAuthenticated, loading } = useAuth();
  const location = useLocation();

  if (loading) {
    // Optional: Replace with a proper loading spinner
    return <PageLoader />;
  }

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return <>{children}</>;
}
