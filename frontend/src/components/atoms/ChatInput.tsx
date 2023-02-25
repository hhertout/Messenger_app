import React, { FormEvent, useContext, useRef } from "react"
import { useParams } from "react-router-dom"
import { sendMessage } from "../../api/Chat"
import { MessageContext } from "../../contexts/MessageContext"
import type { Message } from "../../contexts/MessageContext"
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
      .then(() => {
        target.value = ""
      })
  }
  return (
    <form className="chat-form" onSubmit={handleSubmit}>
      <input className="chat-input" ref={chatInputRef} />
    </form>
  )
}
