import { Routes, Route } from "react-router-dom"
import Home from "./pages/Home"
import Signup from "./pages/Signup"
import "./assets/styles/global.scss"
import Login from "./pages/Login"
import LoggedHome from "./pages/private/LoggedHome"
import Tapbar from "./components/organisms/navbar/Tapbar"
import Invit from "./pages/private/Invit"

function App() {
  return (
    <div className="app">
      <Tapbar />
      <Routes>
        <Route path="/" element={<Home />} />

        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />

        <Route path="/app" element={<LoggedHome />} />
        <Route path="/app/invit" element={<Invit />} />
        
      </Routes>
    </div>
  )
}

export default App
