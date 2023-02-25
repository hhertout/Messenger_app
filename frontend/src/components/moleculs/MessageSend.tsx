import React from "react"
import "./message.scss"

type Props = {
  message: string
}

export default function MessageSend({ message }: Props) {
  return (
    <div className="message-wrapper-sender">
      <div className="message-sent-bubble">{message}</div>
    </div>
  )
}
