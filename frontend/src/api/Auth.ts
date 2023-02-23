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

export const getUserConnected = async () => {
  const route: string = `${import.meta.env.VITE_API_URL}/validate`

  const params: RequestInit = {
    method: "GET",
    headers: {
      "Content-type": "application/json",
    },
    credentials: "include",
  }

  const res = await fetch(route, params)
  if(res.status === 200) {
      const data = await res.json()
      return {user: data.user, status: res.status}
  } else {
    return {user: null, status: res.status }
  }
}

export const logout = async () => {
  const route: string = `${import.meta.env.VITE_API_URL}/logout`

  const params: RequestInit = {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    credentials: "include",
  }

  fetch(route, params)
}
