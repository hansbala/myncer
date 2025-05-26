import { Navigate, useLocation } from "react-router-dom";

function useAuth() {
  // Replace this with your real logic (context, session, cookie, etc.)
  const user = null; // or user object if logged in
  return { user, isAuthenticated: true };
}

export function RequireAuth({ children }: { children: React.ReactNode }) {
  const { isAuthenticated } = useAuth();
  const location = useLocation();

  if (!isAuthenticated) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }

  return <>{children}</>;
}
