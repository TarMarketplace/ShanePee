import type { Meta, StoryObj } from '@storybook/react'
import { fn } from '@storybook/test'

import { UserMenu } from './user-menu'

const meta = {
  title: 'Derivatives/Navbar/UserMenu',
  component: UserMenu,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
  args: {
    user: undefined,
    onLogout: fn(),
  },
} satisfies Meta<typeof UserMenu>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    user: {
      id: '1',
      name: 'จอมน้อย',
    },
  },
  render: ({ user, onLogout }) => (
    <div className='w-fit bg-primary-gradient text-white'>
      <UserMenu user={user} onLogout={onLogout} />
    </div>
  ),
}
