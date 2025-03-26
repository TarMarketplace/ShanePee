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
      last_name: undefined,
      gender: undefined,
      tel: undefined,
      address: {
        house_no: undefined,
        district: undefined,
        province: undefined,
        postcode: undefined,
      },
      payment_method: {
        card_number: undefined,
        card_owner: undefined,
        cvv: undefined,
        expire_date: undefined,
      },
      photo: undefined,
      created_at: '2021-10-01T00:00:00Z',
    },
  },
  render: ({ user, onLogout }) => (
    <div className='w-fit bg-primary-gradient text-white'>
      <UserMenu user={user} onLogout={onLogout} />
    </div>
  ),
}
