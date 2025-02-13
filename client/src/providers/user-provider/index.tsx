'use client'

import { useRouter } from 'next/navigation'
import {
  createContext,
  useCallback,
  useContext,
  useEffect,
  useState,
} from 'react'
import { toast } from 'sonner'

import type { User } from '@/types/users'

import { env } from '@/env'

export interface UserData {
  user: User | null
  setUser: (user: User | null) => void
  fetchUser: () => void
}

const UserContext = createContext<UserData | undefined>(undefined)

function UserProvider({ children }: { children?: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)
  const router = useRouter()

  const fetchUser = useCallback(async () => {
    try {
      const response = await fetch(`${env.NEXT_PUBLIC_BASE_API_URL}/auth/me`, {
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include',
      })

      if (response.ok) {
        const data = await response.json()
        setUser(data)
      } else if (response.status === 401) {
        router.push('/login')
      } else {
        toast.error('Something went wrong')
      }
    } catch {
      toast.error('Something went wrong')
    }
  }, [router])

  useEffect(() => {
    fetchUser()
  }, [fetchUser])

  return (
    <UserContext.Provider
      value={{
        user,
        setUser,
        fetchUser,
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

function useUser(): UserData {
  const context = useContext(UserContext)
  if (context === undefined) {
    throw new Error('useUser must be used within UserProvider')
  }
  return context
}

export { UserContext, UserProvider, useUser }
