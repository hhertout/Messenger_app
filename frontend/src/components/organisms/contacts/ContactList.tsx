import ContactCards from "../../moleculs/ContactCards"
import { v4 as uuid } from "uuid"
import "./contactlist.scss"
import { useContext } from "react"
import { UserContext } from "../../../contexts/UserContext"

export default function ContactList() {
  const { contacts, contactNb } = useContext(UserContext)

  return (
    <div className="contact-wrapper">
      <div>Contact number : {contactNb}</div>
      { contacts.length > 0 ? "No contact found... "
        : contacts.map(contact => {
            return (
                <ContactCards key={uuid()} id={contact.ID} firstname={contact.Firstname} lastname={contact.Lastname} />
            )
          })}
    </div>
  )
}
