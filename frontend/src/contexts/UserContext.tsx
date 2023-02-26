import { createContext, PropsWithChildren, SetStateAction, useEffect, useMemo, useState } from "react"
import { getUserConnected } from "../api/Auth"
import { getContacts } from "../api/Contact"

type User = {
  ID: number
  email: string
  firstname: string
  lastname: string
}

type Contact = {
  ID: number
  Firstname: string
  Lastname: string
}

type UserContextType = {
  currentUser: User | null
  contacts: Contact[]
  contactNb: number
  setCurrentUser: React.Dispatch<SetStateAction<User | null>>
  setContacts: React.Dispatch<SetStateAction<Contact[]>>
  setContactNb: React.Dispatch<SetStateAction<number>>
}

export const UserContext = createContext<UserContextType>({
  currentUser: null,
  contacts: [],
  contactNb: 0,
  setCurrentUser: () => {},
  setContacts: () => {},
  setContactNb: () => {},
})

export const UserContextProvider = ({ children }: PropsWithChildren) => {
  const [currentUser, setCurrentUser] = useState<User | null>(null)
  const [contacts, setContacts] = useState<Contact[]>([])
  const [contactNb, setContactNb] = useState<number>(0)

  useMemo(() => {
    getUserConnected().then(res => {
      if (res.status === 200) {
        setCurrentUser({
          ID: res.user.ID,
          email: res.user.Email,
          firstname: res.user.Firstname,
          lastname: res.user.Lastname,
        })

        getContacts()
          .then(res => res.json())
          .then(data => {
            let contacts = data.contacts
            if(!data.contacts){
                contacts = []
            }
            setContacts(contacts)
            setContactNb(data.contacts_nb)
          })
      }
    })
  }, [])

  useEffect(() => {}, [currentUser])

  return (
    <UserContext.Provider value={{ currentUser, setCurrentUser, contacts, setContacts, contactNb, setContactNb }}>
      {children}
    </UserContext.Provider>
  )
}
