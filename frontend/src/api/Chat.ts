export const getMessages = async (id: string) => {
  const route: string = `${import.meta.env.VITE_API_URL}/message/${id}`

  const params: RequestInit = {
    method: "GET",
    headers: {
      "Content-type": "application/json",
    },
    credentials: "include",
  }

  const res = await fetch(route, params)
  return res
}

export const sendMessage = async (id: string, message: string) => {
  const route: string = `${import.meta.env.VITE_API_URL}/message`

  const body = {
    recipient: parseInt(id),
    message: message,
  }

  const params: RequestInit = {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(body),
  }

  const res = await fetch(route, params)
  return res
}
