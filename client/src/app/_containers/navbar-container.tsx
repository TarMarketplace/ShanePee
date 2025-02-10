'use client'

import { useRouter } from 'next/navigation'
import { useState } from 'react'
import { toast } from 'sonner'

import { Navbar } from '@/components/navbar'

import { useUser } from '@/providers/user-provider'

const NavbarContainer = () => {
  const router = useRouter()
  const { user, setUser } = useUser()
  const [search, setSearch] = useState('')

  const handleLogout = () => {
    setUser(null)
    fetch('/api/auth/logout', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })
      .then(() => {
        toast.success('Successfully logged out')
      })
      .catch(() => {
        toast.error('Something went wrong')
      })
  }

  const handleSearch = () => {
    // TODO: Implement search
    console.log(search)
    router.push(`/search?query=${search}`)
  }

  return (
    <Navbar
      user={user}
      onLogout={handleLogout}
      searchValue={search}
      onChangeSearchValue={setSearch}
      onSearch={handleSearch}
    />
  )
}

export { NavbarContainer }
