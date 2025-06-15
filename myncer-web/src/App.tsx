import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Login } from './pages/Login';
import { RequireAuth } from './RequireAuth';
import { NotFound } from './pages/NotFound';
import { Root } from './layouts/Root';
import { SignUp } from './pages/SignUp';
import { Datasources } from './pages/Datasources';
import { DatasourceAuthPage } from './pages/DatasourceAuthPage';
import { Syncs } from './pages/Syncs';

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
            <Syncs />
          </RequireAuth>
        ),
      },
      {
        index: true,
        path: "syncs",
        element: (
          <RequireAuth>
            <Syncs />
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
            <DatasourceAuthPage datasource='SPOTIFY' />
          </RequireAuth>
        ),
      },
      {
        path: "datasource/youtube/callback",
        element: (
          <RequireAuth>
            <DatasourceAuthPage datasource='YOUTUBE'/>
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
