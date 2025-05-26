import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Home } from './pages/Home';
import { Login } from './pages/Login';
import { RequireAuth } from './RequireAuth';
import { NotFound } from './pages/NotFound';
import { Root } from './layouts/Root';

const router = createBrowserRouter([
  {
    path: "/login",
    element: <Login />,
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
