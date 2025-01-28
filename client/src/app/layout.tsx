import type { Metadata } from 'next'
import { Prompt } from 'next/font/google'

import { Footer } from '@/components/footer'
import { Navbar } from '@/components/navbar'

import { UserProvider } from '@/providers/user-provider'

import { cn } from '@/lib/utils'

import '@/styles/globals.css'

const prompt = Prompt({
  subsets: ['latin', 'thai'],
  variable: '--font-prompt',
  weight: ['400', '600'],
})

// TODO: Add your own metadata
export const metadata: Metadata = {
  title: 'Create Next App',
  description: 'Generated by create next app',
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang='en'>
      <body
        className={cn('min-h-dvh font-prompt antialiased', prompt.variable)}
      >
        <UserProvider>
          <Navbar />
          {children}
          <Footer />
        </UserProvider>
      </body>
    </html>
  )
}
