import React from "react"
import InvitSender from "../../components/moleculs/InvitSender"
import InvitList from "../../components/organisms/contacts/InvitList"
import "./invit.scss"

export default function Invit() {
  return (
    <section className="invitation-list-section">
      <h1>Invitations:</h1>
      <div className="wrapper">
        <h3>Invit your friend</h3>
        <InvitSender />
      </div>
      <div className="wrapper">
        <h3>Your invitations pending : </h3>
        <InvitList />
      </div>
    </section>
  )
}
