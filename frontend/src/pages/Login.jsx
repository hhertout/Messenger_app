import React from "react"
import Loginform from "../components/login/loginform"

export default function Login() {
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
    }
    fetch(route, headers)
      .then(res => {
        res.json()
      })
      .then(data => console.log(data))
      .catch(err => console.error(err))
  }

  return (
    <div>
      <h2>Login</h2>
      <Loginform login={login} />
    </div>
  )
}
