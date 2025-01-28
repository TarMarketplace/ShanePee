import { type VariantProps, cva } from 'class-variance-authority'
import * as React from 'react'

import { cn } from '@/lib/utils'

const buttonCapsuleVariants = cva(
  'inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-full text-base font-semibold transition-colors disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0',
  {
    variants: {
      variant: {
        default: 'bg-primary text-primary-foreground hover:bg-primary/90',
        gradient:
          'bg-primary-gradient text-primary-foreground hover:opacity-90',
      },
      size: {
        default: 'h-10 px-3 py-2',
      },
    },
    defaultVariants: {
      variant: 'default',
      size: 'default',
    },
  }
)

export interface ButtonCapsuleProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonCapsuleVariants> {}

const ButtonCapsule = React.forwardRef<HTMLButtonElement, ButtonCapsuleProps>(
  ({ className, variant, size, ...props }, ref) => {
    return (
      <button
        className={cn(buttonCapsuleVariants({ variant, size }), className)}
        ref={ref}
        {...props}
      />
    )
  }
)
ButtonCapsule.displayName = 'ButtonCapsule'

export { ButtonCapsule, buttonCapsuleVariants }
