export const signup = (email: string, password: string, firstname: string, lastname: string) => {
  const route: string = `${import.meta.env.VITE_API_URL}/signup`

  const logs = {
    email: email,
    password: password,
    firstname: firstname,
    lastname: lastname,
  }

  const param: RequestInit = {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify(logs),
  }

  fetch(route, param)
  .then(res => res.json())
  .then(data => console.log(data))
  .catch(err => console.log(err))
}
