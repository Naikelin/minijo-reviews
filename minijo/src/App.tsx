import React from 'react';
import {
  BrowserRouter as Router,
  Routes,
  Route,
  Link,
  Outlet
} from "react-router-dom";

import ResponsiveAppBar from './components/Navbar'
import './App.css';
import AuthProvider from './context/authContext'
import {useAuth} from './context/authContext'

function App() {

  const {user} = useAuth();

  return (
      <>
        {/* <AuthProvider> */}
          <ResponsiveAppBar />
          <Outlet />
        {/* </AuthProvider> */}
      </>
    );
}

export default App;
