import React, { useContext } from "react"
import { useParams } from "react-router-dom"
import { UserContext } from "../../../contexts/UserContext"
import { FindContact } from "../../../helper/FindContact"
import "./chatheader.scss"

export default function ChatHeader() {
  const { id } = useParams()
  const { contacts } = useContext(UserContext)
  const contact = FindContact(parseInt(id!), contacts)
  return (
    <section className="chat-header">
      {contact.Firstname} {contact.Lastname}
    </section>
  )
}
