import type { Meta, StoryObj } from '@storybook/react'

import { Footer } from '.'

const meta = {
  title: 'Derivatives/Footer',
  component: Footer,
  parameters: {
    layout: 'fullscreen',
  },
  tags: ['autodocs'],
  argTypes: {},
} satisfies Meta<typeof Footer>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {}
