import React, { useContext, useEffect } from "react"
import { Link, useNavigate } from "react-router-dom"
import { UserContext } from "../../../contexts/UserContext"
import "./tapbar.scss"
import { logout } from "../../../api/Auth"

export default function Tapbar() {
  const navigate = useNavigate()
  const { currentUser, setCurrentUser } = useContext(UserContext)

  const handleClick = () => {
    logout().then(() => {
      setCurrentUser(null)
      navigate("/login")
    })
  }

  useEffect(() => {}, [])

  return (
    <nav id="tapbar">
      <div className="links">
        <Link to={"/"}>Home</Link>
        {currentUser !== null ? (
          <>
            <Link to={"/app"}>My space</Link>
            <button onClick={handleClick} className="logout-button">
              Logout
            </button>
          </>
        ) : (
          <Link to={"/login"}>Login</Link>
        )}
      </div>
    </nav>
  )
}
