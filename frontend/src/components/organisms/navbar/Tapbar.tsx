import React, { useContext, useEffect } from "react"
import { Link, useNavigate } from "react-router-dom"
import { UserContext } from "../../../contexts/UserContext"
import "./tapbar.scss"
import { logout } from "../../../api/Auth"
import HomeSvg from "../../atoms/HomeSvg"
import LogoutSvg from "../../atoms/LogoutSvg"
import ContactListSvg from "../../atoms/ContactListSvg"
import LoginSvg from "../../atoms/LoginSvg"

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
        <Link to={"/"}>
          <span className="icon-container">
            <HomeSvg />
            <span className="icon-text">Home</span>
          </span>
        </Link>
        {currentUser !== null ? (
          <>
            <Link to={"/app"}>
              <span className="icon-container">
                <ContactListSvg />
                <span className="icon-text">My space</span>
              </span>
            </Link>
            <button onClick={handleClick} className="logout-button">
              <span className="icon-container">
                <LogoutSvg />
                <span className="icon-text">Logout</span>
              </span>
            </button>
          </>
        ) : (
          <Link to={"/login"}>
            <span className="icon-container">
              <LoginSvg />
              <span className="icon-text">Login</span>
            </span>
          </Link>
        )}
      </div>
    </nav>
  )
}
