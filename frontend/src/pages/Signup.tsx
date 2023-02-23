import SignupForm from "../components/organisms/form/SignupForm"
import { signup } from "../api/Auth"
import { useContext, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import { UserContext } from "../contexts/UserContext"

export default function Signup() {
  const { currentUser } = useContext(UserContext)
  const navigate = useNavigate()
  useEffect(() => {
    if (currentUser !== null) {
      navigate("/")
    }
  }, [currentUser])
  return (
    <div>
      <div className="container-sm">
        <h2>Create your account</h2>

        <SignupForm signup={signup} />
      </div>
    </div>
  )
}
