import { Routes, Route } from "react-router-dom"
import Home from "./pages/Home"
import Signup from "./pages/Signup"
import "./assets/styles/global.scss"
import Login from "./pages/Login"
import LoggedHome from "./pages/private/LoggedHome"

function App() {
  return (
    <div className="app">
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/app" element={<LoggedHome />}></Route>
      </Routes>
    </div>
  )
}

export default App
