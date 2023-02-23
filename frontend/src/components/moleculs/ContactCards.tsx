import React from 'react'
import "./contactCard.scss"

type Props = {
    firstname: string
    lastname: string
}

export default function ContactCards({firstname, lastname}: Props) {
  return (
    <>
      <div>{firstname} {lastname}</div>
    </>
  )
}
