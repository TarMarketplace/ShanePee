import { type VariantProps, cva } from 'class-variance-authority'
import * as React from 'react'

import { cn } from '@/lib/utils'

const buttonVariants = cva(
  'inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-full text-base font-semibold transition-colors disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0',
  {
    variants: {
      variant: {
        filled: 'rounded-sm',
        capsule: 'rounded-full',
      },
      colorVariant: {
        default: 'bg-primary text-primary-foreground hover:bg-primary/90',
        gradient:
          'bg-primary-gradient text-primary-foreground hover:opacity-90',
        outline:
          'border border-primary bg-transparent text-primary hover:bg-secondary-100',
      },
      size: {
        default: 'h-10 px-3 py-2',
      },
    },
    defaultVariants: {
      variant: 'capsule',
      colorVariant: 'default',
      size: 'default',
    },
  }
)

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, colorVariant, variant, size, ...props }, ref) => {
    return (
      <button
        className={cn(
          buttonVariants({ variant, colorVariant, size }),
          className
        )}
        ref={ref}
        {...props}
      />
    )
  }
)
Button.displayName = 'Button'

export { Button, buttonVariants }
