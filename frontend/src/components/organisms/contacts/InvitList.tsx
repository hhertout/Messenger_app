import { useEffect, useState } from "react"
import { getInvits } from "../../../api/Invit"
import "./invitlist.scss"
import { v4 as uuid} from "uuid"
import InvitCard from "../../moleculs/InvitCard"

type Invitation = {
    ID: number
    Firstname: string
    Lastname: string
    Status: string
}

export default function InvitList() {
  const [invitations, setInvitations] = useState<Invitation[]>([])
  const [invitationNb, setInvitationNb] = useState<number>(0)

  useEffect(() => {
    getInvits()
      .then(res => res.json())
      .then(data => {
        if(data.invitations !== null) {
            setInvitationNb(data.invitationsNumber)
            setInvitations(data.invitations)
        } 
      })
  }, [])
  return <>
    {invitations.length === 0 ? <p>You have {invitationNb} invitation</p> : invitations.map(invitation => {
        return (
            <InvitCard key={uuid()} firstname={invitation.Firstname} lastname={invitation.Lastname} id={invitation.ID} />
        )
    })}
  </>
}
