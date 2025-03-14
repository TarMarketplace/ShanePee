import { zodResolver } from '@hookform/resolvers/zod'
import type { Meta, StoryObj } from '@storybook/react'
import { fn } from '@storybook/test'
import { useForm } from 'react-hook-form'
import { z } from 'zod'

import { TrackingInputCard } from '.'

const trackingInputCardSchema = z.object({
  trackingNumberValue: z.string().min(1, 'Tracking number is required'),
  deliveryCompanyValue: z.enum(['Shopee express', 'Kerry', 'Flash']),
})

export type TrackingInputCardSchema = z.infer<typeof trackingInputCardSchema>

const meta: Meta<typeof TrackingInputCard> = {
  title: 'Derivatives/TrackingInputCard',
  component: TrackingInputCard,
  tags: ['autodocs'],
  parameters: {
    layout: 'fullscreen',
  },
}

export default meta
type Story = StoryObj<typeof meta>

const Template = () => {
  const id = 'SHANE-10649'
  const name = 'CRYBABY × Powerpuff Girls Series Figures'

  const form = useForm<TrackingInputCardSchema>({
    resolver: zodResolver(trackingInputCardSchema),
    defaultValues: {
      trackingNumberValue: '',
      deliveryCompanyValue: 'Shopee express',
    },
  })

  return <TrackingInputCard id={id} name={name} form={form} onSubmit={fn()} />
}

export const Default: Story = {
  render: () => <Template />,
}
