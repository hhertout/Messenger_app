import SignupForm from "../components/organisms/form/SignupForm"
import { signup } from "../api/Auth"

export default function Signup() {
  return (
    <div>
      <div className="container-sm">
        <h2>Create your account</h2>

        <SignupForm signup={signup} />
      </div>
    </div>
  )
}
