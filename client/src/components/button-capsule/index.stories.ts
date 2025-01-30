import type { Meta, StoryObj } from '@storybook/react'
import { fn } from '@storybook/test'

import { ButtonCapsule } from '.'

const meta = {
  title: 'Primitives/ButtonCapsule',
  component: ButtonCapsule,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {},
  args: { onClick: fn(), children: 'Button' },
} satisfies Meta<typeof ButtonCapsule>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    variant: 'default',
  },
}

export const Gradient: Story = {
  args: {
    variant: 'gradient',
  },
}
