'use client'

import { useRouter } from 'next/navigation'
import { useState } from 'react'
import { toast } from 'sonner'

import { Navbar } from '@/components/navbar'

import { useUser } from '@/providers/user-provider'

import { logout } from '@/generated/api'

const NavbarContainer = () => {
  const router = useRouter()
  const { user, setUser } = useUser()
  const [search, setSearch] = useState('')

  const handleLogout = async () => {
    const { response } = await logout()

    if (response.ok) {
      setUser(null)
      toast.success('Logged out successfully.')
    } else {
      toast.error('Something went wrong.')
    }
  }

  const handleSearch = () => {
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
