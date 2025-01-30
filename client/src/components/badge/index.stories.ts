import type { Meta, StoryObj } from '@storybook/react'

import { Badge } from '.'

const meta = {
  title: 'Primitives/Badge',
  component: Badge,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {},
  args: { children: 'Badge' },
} satisfies Meta<typeof Badge>

export default meta
type Story = StoryObj<typeof meta>

export const Primary: Story = {
  args: {
    variant: 'primary',
  },
}

export const PrimaryGradient: Story = {
  args: {
    variant: 'primary-gradient',
  },
}

export const Success: Story = {
  args: {
    variant: 'success',
  },
}

export const Info: Story = {
  args: {
    variant: 'info',
  },
}

export const Error: Story = {
  args: {
    variant: 'error',
  },
}

export const Warning: Story = {
  args: {
    variant: 'warning',
  },
}
