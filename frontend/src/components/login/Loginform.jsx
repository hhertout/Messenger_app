import React, { useRef } from "react"

export default function Loginform({login}) {
    const emailInput = useRef()
    const passwordInput = useRef()

  const handleSubmit = e => {
    e.preventDefault()
    const email = emailInput.current.value
    const password = passwordInput.current.value
    login(email, password)
  }

  return (
    <form onSubmit={handleSubmit}>
      <label htmlFor="email">Email</label>
      <input ref={emailInput} id="email" name="email" type="email" />

      <label htmlFor="password">Password</label>
      <input ref={passwordInput} id="password" name="password" type="password" />

      <button>Loggin</button>
    </form>
  )
}
