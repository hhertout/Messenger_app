import LoginForm from "../components/organisms/form/LoginForm"
import { login } from "../api/Auth"
import { useContext, useEffect } from "react"
import { UserContext } from "../contexts/UserContext"
import { useNavigate } from "react-router-dom"

export default function Login() {
  const navigate = useNavigate()
  const { currentUser } = useContext(UserContext)

  useEffect(() => {
    if (currentUser !== null) {
      navigate("/")
    }
  }, [])

  return (
    <div>
      <div className="container-sm">
        <h2>Loggin</h2>
        <LoginForm login={login} />
      </div>
    </div>
  )
}
