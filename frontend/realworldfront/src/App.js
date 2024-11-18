import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import Login from './pages/login';
import Home from './pages/home';
import Nav from './components/nav';

import useToken from './components/useToken'



function App() {

  const { token, setToken } = useToken();
  return (
    <div className="wrapper">

      <Nav />
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/" element={<Home />} />
        </Routes>
      </BrowserRouter >
    </div >
  );
}

export default App;
