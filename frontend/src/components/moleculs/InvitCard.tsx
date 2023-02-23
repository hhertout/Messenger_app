import "./invitcard.scss"

type Props = {
    id: number
    firstname: string
    lastname: string
}

export default function InvitCard({id, firstname, lastname}: Props) {
  return (
    <div>
      {firstname} {lastname}
    </div>
  )
}
