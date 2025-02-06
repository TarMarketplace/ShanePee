import { type VariantProps, cva } from 'class-variance-authority'
import * as React from 'react'

import { Text } from '@/components/text'

import { cn } from '@/lib/utils'

const badgeVariants = cva(
  'inline-flex items-center rounded-sm px-1 py-0.5 text-white shadow-sm',
  {
    variants: {
      variant: {
        primary: 'bg-primary text-primary-foreground',
        'primary-gradient': 'bg-primary-gradient text-primary-foreground',
        success: 'bg-success',
        info: 'bg-info',
        error: 'bg-error',
        warning: 'bg-warning',
      },
    },
    defaultVariants: {
      variant: 'primary',
    },
  }
)

export interface BadgeProps
  extends React.HTMLAttributes<HTMLDivElement>,
    VariantProps<typeof badgeVariants> {}

function Badge({ className, variant, children, ...props }: BadgeProps) {
  return (
    <div className={cn(badgeVariants({ variant }), className)} {...props}>
      <Text variant='sm-regular'>{children}</Text>
    </div>
  )
}

export { Badge, badgeVariants }
