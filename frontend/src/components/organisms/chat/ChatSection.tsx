import { useContext, useEffect, useMemo, useState } from "react"
import { useParams } from "react-router-dom"
import { getMessages } from "../../../api/Chat"
import MessageReceived from "../../moleculs/MessageReceived"
import MessageSend from "../../moleculs/MessageSend"
import { v4 as uuid } from "uuid"
import NoMessage from "../../moleculs/NoMessage"
import "./chatsection.scss"
import { MessageContext } from "../../../contexts/MessageContext"
import { scrollToBot } from "../../../helper/ScrollToBot"

export default function ChatSection() {
  let { id } = useParams<string>()
  const { messages, setMessages } = useContext(MessageContext)

  useEffect(() => {
    if (id) {
      getMessages(id)
        .then(res => res.json())
        .then(data => {
          if (data.messages) {
            const messages = data.messages
            setMessages(messages)
          }
        })
    }
  }, [])

  useEffect(() => {
    scrollToBot("chat-section")
  }, [messages])

  return (
    <section id="chat-section" className="chat-section">
      {messages.length > 0 ? (
        messages.map(message => {
          return (
            <div key={uuid()}>
              {message.Sender ? (
                <MessageSend message={message.Message} />
              ) : (
                <MessageReceived message={message.Message} firstname={message.Firstname} />
              )}
            </div>
          )
        })
      ) : (
        <NoMessage />
      )}
    </section>
  )
}
