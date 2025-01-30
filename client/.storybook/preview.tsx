import type { Preview } from '@storybook/react'
import { Prompt } from 'next/font/google'
import React from 'react'

import '../src/styles/globals.css'

const prompt = Prompt({
  subsets: ['latin', 'thai'],
  variable: '--font-prompt',
  weight: ['400', '600'],
})

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
  decorators: (Story) => (
    <div className={`font-prompt ${prompt.variable}`}>
      <Story />
    </div>
  ),
}

export default preview
