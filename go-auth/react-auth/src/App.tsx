import React from 'react';
import './App.css';
import Login from './Pages/Login';
import Navbar from './components/Navbar';
import { BrowserRouter, Route } from 'react-router-dom';
import Home from './Pages/Home';
import Register from './Pages/Register';


function App() {
  return (
    <div className="App">
      <Navbar/>
      {/* <BrowserRouter>
          <Navbar/>
          <main className="form-signin">
            <Route path="/" element={<Home/>}/>
            <Route path="/login" element={<Login/>}/>
            <Route path="/register" element={<Register/>}/>
        </main>
      </BrowserRouter> */}
    </div>
  );
}

export default App;
