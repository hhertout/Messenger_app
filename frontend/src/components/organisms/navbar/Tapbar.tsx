import React from "react"
import { Link } from "react-router-dom"
import "./tapbar.scss"

export default function Tapbar() {
  return (
    <nav id="tapbar">
      <div className="links">
        <Link to={"/"}>Home</Link>
        <Link to={"/app"}>My space</Link>
      </div>
    </nav>
  )
}
