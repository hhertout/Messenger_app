import React from "react"
import LoginForm from "../components/organisms/form/LoginForm"
import { login } from "../api/Auth"

export default function Login() {
  return (
    <div>
      <div className="container-sm">
        <h2>Loggin</h2>
        <LoginForm login={login} />
      </div>
    </div>
  )
}
