import React, { useEffect, useState } from "react"
import { getContacts } from "../../../api/Contact"
import ContactCards from "../../moleculs/ContactCards"
import {v4 as uuid} from "uuid"
import "./contactlist.scss"

type Contact = {
  Firstname: string
  Lastname: string
}

export default function ContactList() {
  const [contacts, setContacts] = useState<Contact[]>([])
  const [contactNb, setContactNb] = useState<number>(0)

  useEffect(() => {
    getContacts()
      .then(res => res.json())
      .then(data => {
        setContacts(data.contacts)
        setContactNb(data.contacts_nb)
      })
  }, [])

  //useEffect(() => {}, [contacts])

  return (
    <div className="contact-wrapper">
      <div>Nombre de contacts : {contactNb}</div>
      {contacts.length === 0
        ? "No contact found... "
        : contacts.map((contact: Contact) => {
            return (
                <ContactCards key={uuid()} firstname={contact.Firstname} lastname={contact.Lastname} />
            )
          })}
    </div>
  )
}
