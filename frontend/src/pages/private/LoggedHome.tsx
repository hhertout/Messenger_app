import { useNavigate } from "react-router-dom"
import { UserContext } from "../../contexts/UserContext"
import { useContext, useEffect } from "react"
import ContactList from "../../components/organisms/contacts/ContactList"
import UserCard from "../../components/moleculs/UserCard"

export default function LoggedHome() {
  const navigate = useNavigate()
  const { currentUser } = useContext(UserContext)

  useEffect(() => {
    if (currentUser === null) {
      navigate("/")
    }
  }, [currentUser])
  return (
    <main>
      {currentUser === null ? (
        ""
      ) : (
        <>
          <UserCard firstname={currentUser.firstname} lastname={currentUser.lastname} status={"🟢"} />
          <ContactList />
        </>
      )}
    </main>
  )
}
