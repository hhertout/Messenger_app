import { useEffect, useState } from "react"
import "./contactCard.scss"
import { getLastMessage } from "../../api/Contact"
import { Link } from "react-router-dom"

type Props = {
  id: number
  firstname: string
  lastname: string
}

export default function ContactCards({ id, firstname, lastname }: Props) {
  const [lastMessage, setLastMessage] = useState<string>("")
  useEffect(() => {
    getLastMessage(id)
      .then(res => res.json())
      .then(data => setLastMessage(data.message.Message))
  }, [])
  return (
    <Link to={`/app/chat/${id}`}>
    <div className="contact-card">
      <div>
        {firstname} {lastname}
      </div>
      <div className="last-message">{lastMessage === "" ? "No message found" : lastMessage.substring(0, 65) + "..."}</div>
    </div>
    </Link>
  )
}
