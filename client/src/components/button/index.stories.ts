import type { Meta, StoryObj } from '@storybook/react'

import { Button } from '.'

const meta = {
  title: 'Primitives/Button',
  component: Button,
  parameters: {
    layout: 'centered',
  },
  tags: ['autodocs'],
  argTypes: {
    colorVariant: {
      options: ['default', 'gradient'],
      control: {
        type: 'radio',
      },
    },
  },
  args: { children: 'Button' },
} satisfies Meta<typeof Button>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  args: {
    colorVariant: 'default',
  },
}

export const Gradient: Story = {
  args: {
    colorVariant: 'gradient',
  },
}

export const Filled: Story = {
  args: {
    variant: 'filled',
  },
}

export const FilledGradient: Story = {
  args: {
    variant: 'filled',
    colorVariant: 'gradient',
  },
}
