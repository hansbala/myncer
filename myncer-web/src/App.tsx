import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Home } from './pages/Home';
import { Login } from './pages/Login';
import { RequireAuth } from './RequireAuth';
import { NotFound } from './pages/NotFound';
import { Root } from './layouts/Root';
import { SignUp } from './pages/SignUp';
import { Datasources } from './pages/Datasources';
import { DatasourceAuthPage } from './pages/DatasourceAuthPage';

const router = createBrowserRouter([
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/signup",
    element: <SignUp />,
  },
  {
    path: "/",
    element: <Root />,
    children: [
      {
        index: true,
        element: (
          <RequireAuth>
            <Home />
          </RequireAuth>
        ),
      },
      {
        path: "datasources",
        element: (
          <RequireAuth>
            <Datasources />
          </RequireAuth>
        ),
      },
      {
        path: "datasource/spotify/callback",
        element: (
          <RequireAuth>
            <DatasourceAuthPage />
          </RequireAuth>
        ),
      },
    ],
  },
  {
    path: "*",
    element: <NotFound />,
  }
]);


function App() {
  return <RouterProvider router={router} />;
}

export default App
