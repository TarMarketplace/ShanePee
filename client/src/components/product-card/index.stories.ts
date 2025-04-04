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
      owner: {
        first_name: 'Jom',
        last_name: 'Noiz',
        address: {
          district: 'a',
          province: 'a',
          house_no: '1',
          postcode: '10000',
        },
        email: 'a',
        created_at: '2024-08-01T00:00:00Z',
        id: 9007199254740991,
        payment_method: {
          card_number: '1234-5678-9012-3456',
          card_owner: 'Jom Noiz',
          cvv: '123',
          expire_date: '2024-08-01T00:00:00Z',
        },
      },
      owner_id: 1,
      release_date: new Date().toISOString(),
      average_rating: 4.5,
    },
  },
}
