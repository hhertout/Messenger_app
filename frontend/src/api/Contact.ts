export const getContacts = async () => {
  const route: string = `${import.meta.env.VITE_API_URL}/contacts`

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

export const getLastMessage = async (id: number) => {
  const route: string = `${import.meta.env.VITE_API_URL}/lastMessage/${id}`

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
