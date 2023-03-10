import "./usercard.scss"

type Props = {
  firstname: string
  lastname: string
  status: string
}

export default function UserCard({ firstname, lastname, status }: Props) {
  return (
    <div className="usercard">
      <div className="name-wrapper">
        <h1 className="username">
          {firstname} {lastname}
        </h1>
        <div className="userstatus">{status}</div>
      </div>
    </div>
  )
}
