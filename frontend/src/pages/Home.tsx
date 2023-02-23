import { useContext } from "react"
import { UserContext } from "../contexts/UserContext"

export default function Home() {
  const { currentUser } = useContext(UserContext)

  return (
    <div>
      <h1>Hello {currentUser != null ? currentUser.firstname : "world"}</h1>

      
    </div>
  )
}
