import { useEffect } from 'react'
import { BrowserRouter, Route, Routes, Navigate } from 'react-router-dom'
import { AuthPage } from './pages/AuthPage'
import { TaskManagerPage } from './pages/TaskManagerPage'
import { CsvUploadPage } from './pages/CsvUploadPage'
import axios from 'axios'
import { CsrfToken } from './types'
import { CardStatementsPage } from './pages/CardStatementsPage'
import { PublicRoute } from './components/route/PublicRoute'
import { AuthenticatedRoute } from './components/route/AuthenticatedRoute'
import { Layout } from './components/layout/Layout'

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
        {/* 未認証ユーザー向けルート */}
        <Route element={<PublicRoute />}>
          <Route path="/" element={<AuthPage />} />
          <Route path="/login" element={<AuthPage />} />
        </Route>
        
        {/* 認証済みユーザー向けルート */}
        <Route element={<AuthenticatedRoute />}>
          <Route element={<Layout />}>
            <Route path="/task-manager" element={<TaskManagerPage />} />
            <Route path="/csv-upload" element={<CsvUploadPage />} />
            <Route path="/card-statements-page" element={<CardStatementsPage />} />
          </Route>
        </Route>
        
        {/* 存在しないパスへのリダイレクト */}
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App