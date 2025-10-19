import "./App.css";
import Home from "./components/home/Home";
import Header from "./components/header/Header";
import Register from "./components/register/Register";
import Login from "./components/login/Login";
import { Route, Routes, useNavigate } from "react-router-dom";
function App() {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
      </Routes>
    </>
  );
}

export default App;
