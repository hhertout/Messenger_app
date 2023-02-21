import React from "react"
import Loginform from "../components/login/loginform"
import { useNavigate } from "react-router-dom"

export default function Login() {
  const navigate = useNavigate()
  const login = (email, password) => {
    const logs = {
      email: email,
      password: password,
    }
    const route = `${import.meta.env.VITE_API_URL}/login`
    const headers = {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(logs),
      credentials: "include",
    }
    fetch(route, headers)
      .then(res => {
        console.log(res.status)
        if (res.status == 200) {
          navigate("/")
        }
      })
      .catch(err => console.error(err))
  }

  return (
    <div>
      <h2>Login</h2>
      <Loginform login={login} />
    </div>
  )
}
