import type { Meta, StoryObj } from '@storybook/react'

import { ProductCard } from '.'

const meta = {
  title: 'Derivatives/ProductCard',
  component: ProductCard,
  parameters: {
    layout: 'padded',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof ProductCard>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    product: {
      id: 1,
      name: 'Product',
      photo: 'https://placehold.co/250x140.png',
      price: 100,
      availability: true,
      description: 'Description for Product',
      owner_id: 1,
      release_date: new Date().toISOString(),
      average_rating: 4.5,
    },
  },
}
