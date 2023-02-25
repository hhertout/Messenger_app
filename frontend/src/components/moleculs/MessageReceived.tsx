import React from "react"
import "./message.scss"

type Props = {
  message: string
  firstname: string
}

export default function MessageReceived({ message, firstname }: Props) {
  return (
    <div className="message-wrapper-receiver">
      <div className="chat-firstname">{firstname}</div>
      <div className="message-received-bubble">{message}</div>
    </div>
  )
}
