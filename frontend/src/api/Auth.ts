import { useNavigate } from "react-router-dom"

export const signup = async (email: string, password: string, firstname: string, lastname: string) => {
  const route: string = `${import.meta.env.VITE_API_URL}/signup`

  const logs = {
    email: email,
    password: password,
    firstname: firstname,
    lastname: lastname,
  }

  const params: RequestInit = {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify(logs),
  }

  const res = await fetch(route, params)
  return res
}

export const login = async (email: string, password: string) => {
  const route: string = `${import.meta.env.VITE_API_URL}/login`

  const logs = {
    email: email,
    password: password,
  }

  const params: RequestInit = {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify(logs),
    credentials: "include",
  }

  const res = await fetch(route, params)
  return res
}
