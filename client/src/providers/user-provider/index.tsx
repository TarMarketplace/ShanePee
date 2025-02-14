'use client'

import { createContext, useContext, useEffect, useState } from 'react'
import { toast } from 'sonner'

import type { User } from '@/types/users'

import { env } from '@/env'

export interface UserData {
  user: User | null
  setUser: (user: User | null) => void
}

const UserContext = createContext<UserData | undefined>(undefined)

function UserProvider({ children }: { children?: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await fetch(
          `${env.NEXT_PUBLIC_BASE_API_URL}/v1/auth/me`,
          {
            headers: {
              'Content-Type': 'application/json',
            },
            credentials: 'include',
          }
        )

        if (response.ok) {
          const data = await response.json()
          setUser(data)
        }
      } catch {
        toast.error('Something went wrong')
      }
    }

    fetchUser()
  }, [])

  return (
    <UserContext.Provider
      value={{
        user,
        setUser,
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
