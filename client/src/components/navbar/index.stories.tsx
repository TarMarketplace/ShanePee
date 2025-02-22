import type { Meta, StoryObj } from '@storybook/react'
import { fn } from '@storybook/test'
import { useState } from 'react'

import type { User } from '@/generated/api'

import { Navbar } from '.'

const meta = {
  title: 'Derivatives/Navbar',
  component: Navbar,
  // This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/writing-docs/autodocs
  tags: ['autodocs'],
  parameters: {
    // More on how to position stories at: https://storybook.js.org/docs/configure/story-layout
    layout: 'fullscreen',
  },
  args: {
    user: null,
    onLogout: fn(),
    onSearch: fn(),
    onChangeSearchValue: fn(),
    searchValue: '',
  },
} satisfies Meta<typeof Navbar>

export default meta
type Story = StoryObj<typeof meta>

const Template = () => {
  const [user, setUser] = useState<User | null>(null)
  const [searchValue, setSearchValue] = useState('')

  const handleLogout = () => setUser(null)
  const handleSearch = () => console.log(searchValue)
  const handleChangeSearchValue = (value: string) => setSearchValue(value)

  return (
    <Navbar
      user={user}
      onLogout={handleLogout}
      searchValue={searchValue}
      onChangeSearchValue={handleChangeSearchValue}
      onSearch={handleSearch}
    />
  )
}

export const Default: Story = {
  render: () => <Template />,
}
