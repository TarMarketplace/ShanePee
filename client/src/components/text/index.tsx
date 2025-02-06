import { type VariantProps, cva } from 'class-variance-authority'
import React from 'react'

import { cn } from '@/lib/utils'

const textVariants = cva('font-prompt', {
  variants: {
    variant: {
      'heading-5xl': 'text-7xl font-semibold leading-[5rem]',
      'heading-4xl': 'text-[4rem] font-semibold leading-[4.5rem]',
      'heading-3xl': 'text-[3.5rem] font-semibold leading-[4rem]',
      'heading-2xl': 'text-[3rem] font-semibold leading-[3.5rem]',
      'heading-xl': 'text-[2.5rem] font-semibold leading-[3rem]',
      'heading-lg': 'text-[2rem] font-semibold leading-10',
      'heading-md': 'text-[1.5rem] font-semibold leading-8',
      'heading-sm': 'text-[1.25rem] font-semibold leading-7',
      'heading-xs': 'text-[1rem] font-semibold leading-6',
      'heading-xxs': 'text-[0.875rem] font-semibold leading-5',

      'xl-regular': 'text-[1.25rem] font-normal leading-7',
      'xl-semibold': 'text-[1.25rem] font-semibold leading-7',

      'lg-regular': 'text-[1.125rem] font-normal leading-7',
      'lg-semibold': 'text-[1.125rem] font-semibold leading-7',

      'md-regular': 'text-[1rem] font-normal leading-6',
      'md-semibold': 'text-[1rem] font-semibold leading-6',

      'sm-regular': 'text-[0.75rem] font-normal leading-4',
      'sm-semibold': 'text-[0.75rem] font-semibold leading-4',

      'xs-regular': 'text-[0.625rem] font-normal leading-[0.875rem]',
      'xs-semibold': 'text-[0.625rem] font-semibold leading-[0.875rem]',
    },
  },
  defaultVariants: {
    variant: 'md-regular',
  },
})

export interface TextProps
  extends React.HTMLAttributes<HTMLParagraphElement>,
    VariantProps<typeof textVariants> {
  as?: React.ElementType
}

const Text = React.forwardRef<HTMLParagraphElement, TextProps>(
  ({ className, as, variant, ...props }, ref) => {
    let Component: React.ElementType = 'p'

    switch (variant) {
      case 'heading-5xl':
        Component = 'h1'
        break
      case 'heading-4xl':
        Component = 'h2'
        break
      case 'heading-3xl':
        Component = 'h3'
        break
      case 'heading-2xl':
        Component = 'h4'
        break
      case 'heading-xl':
        Component = 'h5'
        break
      case 'heading-lg':
        Component = 'h6'
        break
    }

    if (as) {
      Component = as
    }

    return (
      <Component
        className={cn(textVariants({ variant, className }))}
        ref={ref}
        {...props}
      />
    )
  }
)

Text.displayName = 'Text'

export { Text }
