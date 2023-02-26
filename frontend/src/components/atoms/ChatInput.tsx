import React, { FormEvent, useContext, useRef } from "react"
import { useParams } from "react-router-dom"
import { sendMessage } from "../../api/Chat"
import { MessageContext } from "../../contexts/MessageContext"
import "./chatinput.scss"

export default function ChatInput() {
  const { id } = useParams()
  const { messages, setMessages } = useContext(MessageContext)
  const chatInputRef = useRef<HTMLInputElement>(null)

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()

    const target = chatInputRef.current!

    sendMessage(id!, target.value)
      .then(res => res.json())
      .then(data => {
        let newMessage = data.message
        newMessage = { ...newMessage, Sender: true }
        const message = [...messages, newMessage]
        setMessages(message)

        target.value = ""
      })
  }
  return (
    <form className="chat-form" onSubmit={handleSubmit}>
      <input className="chat-input" ref={chatInputRef} placeholder={"Send your message"}/>
    </form>
  )
}
