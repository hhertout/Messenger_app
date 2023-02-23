import { createContext, PropsWithChildren, SetStateAction, useEffect, useState } from "react"
import { getUserConnected } from "../api/Auth"

type User = {
  ID: number
  email: string
  firstname: string
  lastname: string
}

type UserContextType = {
  currentUser: User | null
  setCurrentUser: React.Dispatch<SetStateAction<User | null>>
}

export const UserContext = createContext<UserContextType>({
  currentUser: null,
  setCurrentUser: () => {},
})

export const UserContextProvider = ({ children }: PropsWithChildren) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null)

  useEffect(() => {
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
  }, [])

  useEffect(() => {}, [currentUser])

  return <UserContext.Provider value={{ currentUser, setCurrentUser }}>{children}</UserContext.Provider>
}
