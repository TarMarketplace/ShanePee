import { type VariantProps, cva } from 'class-variance-authority'
import React from 'react'

import { cn } from '@/lib/utils'

const textVariants = cva('font-prompt', {
  variants: {
    desktopVariant: {
      'heading-5xl': 'md:text-[7rem] md:font-semibold md:leading-[5rem]',
      'heading-4xl': 'md:text-[4rem] md:font-semibold md:leading-[4.5rem]',
      'heading-3xl': 'md:text-[3.5rem] md:font-semibold md:leading-[4rem]',
      'heading-2xl': 'md:text-[3rem] md:font-semibold md:leading-[3.5rem]',
      'heading-xl': 'md:text-[2.5rem] md:font-semibold md:leading-[3rem]',
      'heading-lg': 'md:text-[2rem] md:font-semibold md:leading-10',
      'heading-md': 'md:text-[1.5rem] md:font-semibold md:leading-8',
      'heading-sm': 'md:text-[1.25rem] md:font-semibold md:leading-7',
      'heading-xs': 'md:text-[1rem] md:font-semibold md:leading-6',
      'heading-xxs': 'md:text-[0.875rem] md:font-semibold md:leading-5',

      'xl-regular': 'md:text-[1.25rem] md:font-normal md:leading-7',
      'xl-semibold': 'md:text-[1.25rem] md:font-semibold md:leading-7',

      'lg-regular': 'md:text-[1.125rem] md:font-normal md:leading-7',
      'lg-semibold': 'md:text-[1.125rem] md:font-semibold md:leading-7',

      'md-regular': 'md:text-[1rem] md:font-normal md:leading-6',
      'md-semibold': 'md:text-[1rem] md:font-semibold md:leading-6',

      'sm-regular': 'md:text-[0.75rem] md:font-normal md:leading-4',
      'sm-semibold': 'md:text-[0.75rem] md:font-semibold md:leading-4',

      'xs-regular': 'md:text-[0.625rem] md:font-normal md:leading-[0.875rem]',
      'xs-semibold':
        'md:text-[0.625rem] md:font-semibold md:leading-[0.875rem]',
    },
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

      'sm-regular': 'text-[0.875rem] font-normal leading-[1.125rem]',
      'sm-semibold': 'text-[0.875rem] font-semibold leading-[1.125rem]',

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
  ({ className, as, desktopVariant, variant, ...props }, ref) => {
    let Component: React.ElementType = 'p'

    switch (variant) {
      case 'heading-5xl':
        Component = 'h1'
        break
      case 'heading-4xl':
      case 'heading-3xl':
        Component = 'h2'
        break
      case 'heading-2xl':
        Component = 'h3'
        break
      case 'heading-xl':
      case 'heading-lg':
        Component = 'h4'
        break
      case 'heading-md':
      case 'heading-sm':
        Component = 'h5'
        break
      case 'heading-xxs':
      case 'heading-xs':
        Component = 'h6'
        break
    }

    if (as) {
      Component = as
    }

    return (
      <Component
        className={cn(textVariants({ variant, desktopVariant, className }))}
        ref={ref}
        {...props}
      />
    )
  }
)

Text.displayName = 'Text'

export { Text }
