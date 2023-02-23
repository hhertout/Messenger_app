import React, { useContext, useRef, useState } from "react"
import { useNavigate } from "react-router-dom"
import Button from "../../atoms/Button"
import { FormRaw } from "../../moleculs/FormRaw"
import { getUserConnected } from "../../../api/Auth"
import { UserContext } from "../../../contexts/UserContext"

type Props = {
  login: (email: string, password: string) => Promise<Response>
}

export default function LoginForm({ login }: Props) {
  const navigate = useNavigate()
  const { setCurrentUser } = useContext(UserContext)
  const emailInput = useRef<HTMLInputElement>(null!)
  const passwordInput = useRef<HTMLInputElement>(null!)
  const [error, setError] = useState<string>("")

  const handleSubmit = async (e: React.SyntheticEvent) => {
    e.preventDefault()
    const email = emailInput.current.value
    const password = passwordInput.current.value

    const res = await login(email, password)
    if (res.status != 200) {
      console.error("Error")
      setError("incorrect login or password")
    } else {
      getUserConnected().then(res => {
        if (res.status === 200) {
          setCurrentUser({
            ID: res.user.ID,
            email: res.user.Email,
            firstname: res.user.Firstname,
            lastname: res.user.Lastname,
          })
        }
      })
      navigate("/app")
    }
  }

  return (
    <form onSubmit={handleSubmit}>
      <FormRaw name={"email"} title={"Email"} type={"email"} isRequired={true} inputRef={emailInput} />
      <FormRaw name={"password"} title={"Password"} type={"password"} isRequired={true} inputRef={passwordInput} />

      <p className="error">{error}</p>

      <Button title={"Loggin"} />
    </form>
  )
}
