import React, { useContext, useEffect } from "react"
import { Link } from "react-router-dom"
import { UserContext } from "../../../contexts/UserContext"
import "./tapbar.scss"

export default function Tapbar() {
    const { currentUser } = useContext(UserContext)

    useEffect(() => {}, [])

  return (
    <nav id="tapbar">
      <div className="links">
        <Link to={"/"}>Home</Link>
        { currentUser !== null ? <Link to={"/app"}>My space</Link> : <Link to={"/login"}>Login</Link>}
        
      </div>
    </nav>
  )
}
