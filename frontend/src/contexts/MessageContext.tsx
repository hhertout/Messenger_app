import { createContext, PropsWithChildren, SetStateAction, useState } from "react"

export type Message = {
  ID: number
  CreatedAt: string
  Firstname: string
  Sender: boolean
  Message: string
}

type MessageContextType = {
    messages: Message[]
    setMessages : React.Dispatch<SetStateAction<Message[]>>
}

export const MessageContext = createContext<MessageContextType>({
    messages: [{ID: 1, CreatedAt: "", Firstname: "", Sender: true, Message: ""}], 
    setMessages: () => {}
})

export const MessageContextProvider = ({ children }: PropsWithChildren) => {
    const [messages, setMessages] = useState<Message[]>([])

  return <MessageContext.Provider value={{ messages, setMessages }}>{children}</MessageContext.Provider>
}
