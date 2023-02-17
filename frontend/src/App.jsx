import { Route, Routes } from 'react-router-dom'
import './App.css'
import Home from './pages/home'

function App() {
  return (
    <div className="App">
        <Routes>
            <Route path='/' element={<Home />} />
            <Route path='/login' element={<Login />} />
        </Routes>
    </div>
  )
}

export default App
