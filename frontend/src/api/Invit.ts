export const getInvits = async () => {
    const route: string = `${import.meta.env.VITE_API_URL}/invite/pending`
  
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

export const acceptInvit = async(id: number) => {
    const route: string = `${import.meta.env.VITE_API_URL}/invite/${id}`
  
    const params: RequestInit = {
      method: "PUT",
      headers: {
        "Content-type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({status: "accepted"})
    }
  
    const res = await fetch(route, params)
    return res
}

export const rejectInvit = async(id: number) => {
    const route: string = `${import.meta.env.VITE_API_URL}/invite/${id}`
  
    const params: RequestInit = {
      method: "DELETE",
      headers: {
        "Content-type": "application/json",
      },
      credentials: "include",
    }
  
    const res = await fetch(route, params)
    return res
}

export const sendInvit = async(email: string) => {
    const route: string = `${import.meta.env.VITE_API_URL}/invite`
    const body = {
        email: email
    }
    const params: RequestInit = {
      method: "POST",
      headers: {
        "Content-type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(body)
    }
  
    const res = await fetch(route, params)
    return res
}