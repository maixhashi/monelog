import { useMutation } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { login, register, logout, UserCredential } from '../../api/auth'

export const useMutateAuth = () => {
  const navigate = useNavigate()

  const loginMutation = useMutation(
    (credential: UserCredential) => login(credential),
    {
      onSuccess: () => {
        navigate('/task-manager')
      },
      onError: (err: any) => {
        alert(`Login failed: ${err.message}`)
      },
    }
  )

  const registerMutation = useMutation(
    (credential: UserCredential) => register(credential),
    {
      onError: (err: any) => {
        alert(`Registration failed: ${err.message}`)
      },
    }
  )

  const logoutMutation = useMutation(() => logout(), {
    onSuccess: () => {
      navigate('/')
    },
    onError: (err: any) => {
      alert(`Logout failed: ${err.message}`)
    },
  })

  return { loginMutation, registerMutation, logoutMutation }
}