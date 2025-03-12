import type { Meta, StoryObj } from '@storybook/react'

import { OrderCard } from '.'

const meta = {
  title: 'Derivatives/OrderCard',
  component: OrderCard,
  parameters: {
    layout: 'padded',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof OrderCard>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    buyer_name: 'JomNoiz',
    order: {
      buyer_id: 1234,
      created_at: '2024-08-01T00:00:00Z',
      delivery_service: 'Shopee express',
      id: 97,
      seller_id: 4321,
      status: 'PREPARING',
      tracking_number: 'SPX123456789',
    },
    order_items: [
      {
        $schema: 'https://example.com/schemas/ArtToy.json',
        availability: true,
        description: 'CRYBABY × Powerpuff Girls Series Figures',
        id: 9007199254740991,
        name: 'CRYBABY × Powerpuff Girls Series Figures',
        owner_id: 9007199254740991,
        photo: 'https://placehold.co/150x96.png',
        price: 100,
        release_date: '2021-01-01T00:00:00Z',
      },
      {
        $schema: 'https://example.com/schemas/ArtToy.json',
        availability: true,
        description: 'CRYBABY × Powerpuff Girls Series Figures',
        id: 9007199254740991,
        name: 'CRYBABY × Powerpuff Girls Series Figures',
        owner_id: 9007199254740991,
        photo: 'https://placehold.co/150x96.png',
        price: 100,
        release_date: '2021-01-01T00:00:00Z',
      },
      {
        $schema: 'https://example.com/schemas/ArtToy.json',
        availability: true,
        description: 'CRYBABY × Powerpuff Girls Series Figures',
        id: 9007199254740991,
        name: 'CRYBABY × Powerpuff Girls Series Figures',
        owner_id: 9007199254740991,
        photo: 'https://placehold.co/150x96.png',
        price: 100,
        release_date: '2021-01-01T00:00:00Z',
      },
    ],
  },
}
