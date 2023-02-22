import React from 'react'
import LoginForm from '../components/organisms/form/LoginForm'
import { login } from '../api/Auth'

export default function Login() {
  return (
    <div>
        <LoginForm login={login} />
    </div>
  )
}
