import type { Meta, StoryObj } from '@storybook/react'

import { OrderStatus } from './order-status'

const meta = {
  title: 'Derivatives/OrderStatus',
  component: OrderStatus,
  parameters: {
    layout: 'padded',
  },
  tags: ['autodocs'],
  argTypes: {
    status: {
      control: { type: 'select' },
      options: ['PENDING', 'SHIPPING', 'COMPLETED'],
    },
  },
} satisfies Meta<typeof OrderStatus>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    status: 'PENDING',
  },
}
