import React, { useRef, useState } from "react"
import Button from "../../atoms/Button"
import { FormRaw } from "../../moleculs/FormRaw"

type Props = {
  signup: Function
}

export default function SignupForm({ signup }: Props) {
  const [password, setPassword] = useState<string>("")
  const [confirmPassword, setConfirmPassword] = useState<string>("")
  const [error, SetError] = useState<string>("")
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
      />
      <FormRaw
        type={"password"}
        name={"confirmPassword"}
        title={"Confirm your Password"}
        onChange={handleConfirmPwdChange}
        value={confirmPassword}
        isRequired={true}
      />
      {password != confirmPassword ? <p className="error">{error}</p> : null}
      <FormRaw type={"firstname"} name={"firstname"} title={"Firstname"} inputRef={firstnameInput} isRequired={true} />
      <FormRaw type={"lastname"} name={"lastname"} title={"Lastname"} inputRef={lastnameInput} isRequired={true} />
      <Button title={"Create my account"} />
    </form>
  )
}
