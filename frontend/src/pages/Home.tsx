import { useContext } from "react"
import { UserContext } from "../contexts/UserContext"
import "./home.scss"

export default function Home() {
  const { currentUser } = useContext(UserContext)

  return (
    <div>
      <section className="home-section-title">
        <div className="container">
          <h1>Messenger App</h1>
        </div>
      </section>
      <section className="hero-section-wrapper">
        <div className="container hero-section-container">
          <h2>
            Hello <span className="name">{currentUser != null ? currentUser.firstname : ""}</span> !
          </h2>
        </div>
      </section>
    </div>
  )
}
