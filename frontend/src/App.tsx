import { useEffect } from 'react'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { AuthPage } from './pages/AuthPage'
import { TaskManagerPage } from './pages/TaskManagerPage'
import axios from 'axios'
import { CsrfToken } from './types'

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true
    const getCsrfToken = async () => {
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf-token`
      )
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token
    }
    getCsrfToken()
  }, [])
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<AuthPage />} />
        <Route path="/task-manager" element={<TaskManagerPage />} />
      </Routes>
    </BrowserRouter>
  )
}
export default App