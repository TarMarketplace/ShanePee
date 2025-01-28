'use client'

import { createContext, useContext, useState } from 'react'

import { type User } from '@/types/user'

export interface UserData {
  user: User | null
  setUser: (user: User | null) => void
}

const UserContext = createContext<UserData | undefined>(undefined)

function UserProvider({ children }: { children?: React.ReactNode }) {
  const [user, setUser] = useState<User | null>(null)

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
