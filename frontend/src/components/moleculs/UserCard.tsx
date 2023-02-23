import React from "react"

type Props = {
  firstname: string
  lastname: string
  status: string
}

export default function UserCard({ firstname, lastname, status }: Props) {
  return (
    <section>
      <h1>
        {firstname} {lastname}
      </h1>
      <div>{status}</div>
    </section>
  )
}
