import React, { useContext, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { logout } from "../api/Auth"
import { UserContext } from "../contexts/UserContext"

export default function Home() {
  const { currentUser, setCurrentUser } = useContext(UserContext)
  const navigate = useNavigate()

  const handleClick = () => {
    logout().then(() => {
      setCurrentUser(null)
      navigate("/login")
    })
  }
  
  return (
    <div>
      <h1>Hello {currentUser != null ? currentUser.firstname : "world"}</h1>

      <button onClick={handleClick}>Loggout</button>
    </div>
  )
}
