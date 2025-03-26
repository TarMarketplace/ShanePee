import { type VariantProps, cva } from 'class-variance-authority'
import * as React from 'react'

import { Text } from '@/components/text'

import { cn } from '@/lib/utils'

const badgeVariants = cva(
  'inline-flex items-center rounded-sm px-1 py-0.5 text-white',
  {
    variants: {
      variant: {
        primary: 'bg-primary text-primary-foreground shadow-sm',
        'primary-gradient':
          'bg-primary-gradient text-primary-foreground shadow-sm',
        success:
          'rounded-lg border border-success-400 bg-success-50 text-success-400',
        info: 'rounded-lg border border-info-400 bg-info-50 text-info-400',
        error: 'bg-error shadow-sm',
        warning:
          'rounded-lg border border-warning-500 bg-warning-50 text-warning-500',
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
