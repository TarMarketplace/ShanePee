import { type VariantProps, cva } from 'class-variance-authority'
import React from 'react'

import { cn } from '@/lib/utils'

const textVariants = cva('font-prompt', {
  variants: {
    variant: {
      'heading-5xl':
        'text-7xl font-semibold leading-[5rem] tracking-[-0.0625em]',
      'heading-4xl':
        'text-[4rem] font-semibold leading-[4.5rem] tracking-[-0.0625em]',
      'heading-3xl':
        'text-[3.5rem] font-semibold leading-[4rem] tracking-[-0.0625em]',
      'heading-2xl':
        'text-[3rem] font-semibold leading-[3.5rem] tracking-[-0.0625em]',
      'heading-xl':
        'text-[2.5rem] font-semibold leading-[3rem] tracking-[-0.0625em]',
      'heading-lg': 'text-[2rem] font-semibold leading-10 tracking-[-0.0625em]',
      'heading-md':
        'text-[1.5rem] font-semibold leading-8 tracking-[-0.03125em]',
      'heading-sm':
        'text-[1.25rem] font-semibold leading-7 tracking-[-0.03125em]',
      'heading-xs': 'text-[1rem] font-semibold leading-6 tracking-[-0.03125em]',
      'heading-xxs': 'text-[0.875rem] font-semibold leading-5 tracking-normal',

      'xl-regular': 'text-[1.25rem] font-normal leading-7 tracking-normal',
      'xl-semibold': 'text-[1.25rem] font-semibold leading-7 tracking-normal',

      'lg-regular': 'text-[1.125rem] font-normal leading-7 tracking-normal',
      'lg-semibold': 'text-[1.125rem] font-semibold leading-7 tracking-normal',

      'md-regular': 'text-[1rem] font-normal leading-6 tracking-normal',
      'md-semibold': 'text-[1rem] font-semibold leading-6 tracking-normal',

      'sm-regular': 'text-[0.75rem] font-normal leading-4 tracking-normal',
      'sm-semibold': 'text-[0.75rem] font-semibold leading-4 tracking-normal',

      'xs-regular':
        'text-[0.625rem] font-normal leading-[0.875rem] tracking-normal',
      'xs-semibold':
        'text-[0.625rem] font-semibold leading-[0.875rem] tracking-normal',
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
    let Component: React.ElementType = as ?? 'p'

    if (!as) {
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
