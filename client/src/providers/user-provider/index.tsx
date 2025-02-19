'use client'

import { createContext, useContext, useEffect, useState } from 'react'

import type { User } from '@/generated/api'
import { me } from '@/generated/api'

export interface UserData {
  user: User | null
  setUser: (user: User | null) => void
  fetchUser: () => void
}

const UserContext = createContext<UserData | undefined>(undefined)

function UserProvider({ children }: { children?: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)

  const fetchUser = async () => {
    try {
      const { response, data } = await me()
      if (!data) {
        return
      }
      if (response.ok) {
        setUser(data)
      }
    } catch {
      setUser(null)
    }
  }

  useEffect(() => {
    if (!user) fetchUser()
  }, [user])

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
