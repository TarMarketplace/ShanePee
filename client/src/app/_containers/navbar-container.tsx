'use client'

import { useRouter } from 'next/navigation'
import { useState } from 'react'

import { Navbar } from '@/components/navbar'

import { useUser } from '@/providers/user-provider'

const NavbarContainer = () => {
  const router = useRouter()
  const { user, setUser } = useUser()
  const [search, setSearch] = useState('')

  const handleLogin = () => {
    router.push('/login')
    // setUser({ id: '1', name: 'lnwJoZaSodaSing+' })
  }

  const handleLogout = () => {
    setUser(null)
  }

  const handleSearch = () => {
    // TODO: Implement search
    console.log(search)
    router.push(`/search?query=${search}`)
  }

  return (
    <Navbar
      user={user}
      onLogin={handleLogin}
      onLogout={handleLogout}
      searchValue={search}
      onChangeSearchValue={setSearch}
      onSearch={handleSearch}
    />
  )
}

export { NavbarContainer }
