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
      id: 1,
      email: 'jomnoiz@example.com',
      first_name: 'จอมน้อย',
      last_name: null,
      gender: null,
      tel: null,
      address: {
        house_no: null,
        district: null,
        province: null,
        postcode: null,
      },
      payment_method: {
        card_number: null,
        card_owner: null,
        cvv: null,
        expire_date: null,
      },
      photo: null,
    },
  },
  render: ({ user, onLogout }) => (
    <div className='w-fit bg-primary-gradient text-white'>
      <UserMenu user={user} onLogout={onLogout} />
    </div>
  ),
}
