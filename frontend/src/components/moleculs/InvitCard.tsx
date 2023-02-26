import { SyntheticEvent } from "react"
import { acceptInvit, rejectInvit } from "../../api/Invit"
import "./invitcard.scss"

type Props = {
  id: number
  firstname: string
  lastname: string
}

export default function InvitCard({ id, firstname, lastname }: Props) {
  const handleAccept = (e: SyntheticEvent) => {
    e.preventDefault()
    acceptInvit(id)
  }
  const handleReject = (e: SyntheticEvent) => {
    e.preventDefault()
    rejectInvit(id)
  }
  return (
    <div className="invitcard-wrapper">
      <div className="name">
        {firstname} {lastname}
      </div>
      <div>
        <button className="accept" onClick={handleAccept}>
          ✅
        </button>
        <button className="reject" onClick={handleReject}>
          🚫
        </button>
      </div>
    </div>
  )
}
