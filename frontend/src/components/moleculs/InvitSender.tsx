import React, { FormEvent, useRef, useState } from "react"
import { sendInvit } from "../../api/Invit"
import "./invitsender.scss"

export default function InvitSender() {
  const [feedBackMessage, setFeedBackMessage] = useState("")
  const inputRef = useRef<HTMLInputElement>(null)

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault()
    const target = inputRef.current!

    sendInvit(target.value)
      .then(res => res.json())
      .then(data => {
        console.log(data)
        setFeedBackMessage(data.message)
        if(data.message === "invitation sent") {
            target.value = ""
        }
      })
  }

  return (
    <div>
      <p>{feedBackMessage}</p>
      <form onSubmit={handleSubmit}>
        <input type={"email"} ref={inputRef} />
        <button type="submit">Send</button>
      </form>
    </div>
  )
}
