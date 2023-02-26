import React from 'react'
import InvitSender from '../../components/moleculs/InvitSender'
import InvitList from '../../components/organisms/contacts/InvitList'

export default function Invit() {
  return (
    <section>
        <InvitSender />
      Invitations:
        <InvitList />
    </section>
  )
}
