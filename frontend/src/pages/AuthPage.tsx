import { useState, FormEvent } from 'react'
import { ComputerDesktopIcon, ArrowPathIcon } from '@heroicons/react/24/solid'
import { useMutateAuth } from '../hooks/mutateHooks/useMutateAuth'
// 新しいAPIをインポート
import { UserCredential } from '../api/auth'
import '../styles/pages/auth-page/AuthPage.css'

export const AuthPage = () => {
  const [email, setEmail] = useState('')
  const [pw, setPw] = useState('')
  const [isLogin, setIsLogin] = useState(true)
  const { loginMutation, registerMutation } = useMutateAuth()

  const submitAuthHandler = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    // 新しいUserCredentialオブジェクトを作成
    const credential: UserCredential = { email, password: pw }
    
    if (isLogin) {
      loginMutation.mutate(credential)
    } else {
      await registerMutation
        .mutateAsync(credential)
        .then(() => loginMutation.mutate(credential))
    }
  }

  return (
    <div className="auth-container">
      <div className="auth-header">
        <ComputerDesktopIcon className="auth-header-icon" />
        <span className="auth-title">my tech blog</span>
        <ComputerDesktopIcon className="auth-header-icon" />
      </div>
      <h2>{isLogin ? 'Login' : 'Create a new account'}</h2>
      <form className="auth-form" onSubmit={submitAuthHandler}>
        <input
          className="auth-input"
          name="email"
          type="email"
          autoFocus
          placeholder="Email address"
          onChange={(e) => setEmail(e.target.value)}
          value={email}
        />
        <input
          className="auth-input"
          name="password"
          type="password"
          placeholder="Password"
          onChange={(e) => setPw(e.target.value)}
          value={pw}
        />
        <button
          className="auth-button"
          disabled={!email || !pw}
          type="submit"
        >
          {isLogin ? 'Login' : 'Sign Up'}
        </button>
      </form>
      <ArrowPathIcon
        onClick={() => setIsLogin(!isLogin)}
        className="auth-switch"
      />
    </div>
  )
}