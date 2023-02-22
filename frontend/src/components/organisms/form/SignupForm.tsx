import React, { useRef, useState } from "react"
import Button from "../../atoms/Button"
import { FormRaw } from "../../moleculs/FormRaw"

type Props = {
  signup: (email: string, password: string, firstname:string, lastname:string) => Promise<Response>
}

export default function SignupForm({ signup }: Props) {
  const [password, setPassword] = useState<string>("")
  const [confirmPassword, setConfirmPassword] = useState<string>("")
  const [error, setError] = useState<string | null >(null)
  const emailInput = useRef<HTMLInputElement>(null!)
  const firstnameInput = useRef<HTMLInputElement>(null!)
  const lastnameInput = useRef<HTMLInputElement>(null!)

  const handleSubmit = (e: React.SyntheticEvent) => {
    e.preventDefault()
    const email = emailInput.current.value
    const firstname = firstnameInput.current.value
    const lastname = lastnameInput.current.value

    signup(email, password, firstname, lastname)
  }

  const handlePwdChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.currentTarget.value)
  }

  const handleConfirmPwdChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setError(null)
    setConfirmPassword(e.currentTarget.value)
  }

  return (
    <form onSubmit={handleSubmit}>
      <FormRaw type={"email"} name={"email"} title={"Email"} inputRef={emailInput} />
      <FormRaw
        type={"password"}
        name={"password"}
        title={"Password"}
        onChange={handlePwdChange}
        value={password}
        isRequired={true}
        isValid={password === confirmPassword}
      />
      <FormRaw
        type={"password"}
        name={"confirmPassword"}
        title={"Confirm your Password"}
        onChange={handleConfirmPwdChange}
        value={confirmPassword}
        isRequired={true}
        isValid={password === confirmPassword}
      />
      {password != confirmPassword ? <p className="error">⚠️ Passwords must be equals</p> : null}
      <FormRaw type={"firstname"} name={"firstname"} title={"Firstname"} inputRef={firstnameInput} isRequired={true} />
      <FormRaw type={"lastname"} name={"lastname"} title={"Lastname"} inputRef={lastnameInput} isRequired={true} />
      <Button title={"Create my account"} color={"primary"}/>
    </form>
  )
}
