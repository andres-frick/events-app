import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Login from './views/Login';
import reportWebVitals from './reportWebVitals';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import EventList from './views/EventList';
import EventEdit from './views/EventEdit';
import RegisterList from './views/RegisterList';
import EventView from './views/EventView';
import EventNew from './views/EventNew';

const root = ReactDOM.createRoot(document.getElementById('root'));

const router = createBrowserRouter([
  {
    path: "/",
    element: <Login />,
  },
  {
    path: "/app",
    element: <EventList />,
  },
  {
    path: "/app/event/:id",
    element: <EventEdit />,
  },
  {
    path: "/event/registrations",
    element: <RegisterList />,
  },
  {
    path: "/event/:id",
    element: <EventView />,
  },
  {
    path: "/event/new",
    element: <EventNew />,
  }

]);


root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
