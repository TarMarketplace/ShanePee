import type { Meta, StoryObj } from '@storybook/react'

import { Text } from '.'

const meta = {
  title: 'Primitives/Text',
  component: Text,
  parameters: {
    layout: 'padded',
  },
  tags: ['autodocs'],
} satisfies Meta<typeof Text>

export default meta
type Story = StoryObj<typeof meta>

export const Default: Story = {
  argTypes: {
    variant: {
      control: {
        type: 'select',
        options: [
          'heading-5xl',
          'heading-4xl',
          'heading-3xl',
          'heading-2xl',
          'heading-xl',
          'heading-lg',
          'heading-md',
          'heading-sm',
          'heading-xs',
          'lg-regular',
          'lg-semibold',
          'md-regular',
          'md-semibold',
          'sm-regular',
          'sm-semibold',
          'xs-regular',
          'xs-semibold',
        ],
      },
    },
  },
  args: {
    children: 'Text',
    as: 'p',
    variant: 'md-regular',
  },
}

export const Showcase: Story = {
  render: () => {
    return (
      <div className='flex w-full flex-col gap-2'>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-5xl'>Heading 5xl</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 4.5rem(72px)
            <br /> font weight: 600
            <br /> line height: 5rem(80px)
            <br /> letter spacing: -0.0625em(-1px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-4xl'>Heading 4xl</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 3rem(48px)
            <br /> font weight: 600
            <br /> line height: 3.5rem(56px)
            <br />
            letter spacing: -0.0625em(-1px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-3xl'>Heading 3xl</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 2.25rem(36px)
            <br /> font weight: 600
            <br /> line height: 2.5rem(40px)
            <br /> letter spacing: -0.0625em(-1px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-2xl'>Heading 2xl</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1.875rem(30px)
            <br /> font weight: 600
            <br /> line height: 2.25rem(36px)
            <br /> letter spacing: -0.0625em(-1px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-xl'>Heading xl</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1.5rem(24px)
            <br /> font weight: 600
            <br /> line height: 1.75rem(28px)
            <br /> letter spacing: -0.0625em(-1px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-lg'>Heading lg</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1.25rem(20px)
            <br /> font weight: 600
            <br /> line height: 1.25rem(20px)
            <br /> letter spacing: -0.0625em(-1px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-md'>Heading md</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1rem(16px)
            <br /> font weight: 600
            <br /> line height: 1rem(16px)
            <br />
            letter spacing: -0.03125em(-0.5px)
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-sm'>Heading sm</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 0.875rem(14px)
            <br /> font weight: 600
            <br /> line height: 0.875rem(14px)
            <br /> letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='heading-xs'>Heading xs</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 0.75rem(12px)
            <br /> font weight: 600
            <br /> line height: 0.75rem(12px)
            <br /> letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='lg-regular'>Lg Regular</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1.125rem(18px)
            <br /> font weight: 400
            <br /> line height: 1.25rem(20px)
            <br /> letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='lg-semibold'>Lg Semibold</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1.125rem(18px)
            <br /> font weight: 600
            <br /> line height: 1.25rem(20px)
            <br /> letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='md-regular'>Md Regular</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1rem(16px)
            <br /> font weight: 400
            <br /> line height: 1rem(16px)
            <br />
            letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='md-semibold'>Md Semibold</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 1rem(16px)
            <br /> font weight: 600
            <br /> line height: 1rem(16px)
            <br />
            letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='sm-regular'>Sm Regular</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 0.875rem(14px)
            <br /> font weight: 400
            <br /> line height: 0.875rem(14px)
            <br /> letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='sm-semibold'>Sm Semibold</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 0.875rem(14px)
            <br /> font weight: 600
            <br /> line height: 0.875rem(14px)
            <br /> letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='xs-regular'>Xs Regular</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 0.75rem(12px)
            <br /> font weight: 400
            <br /> line height: 1rem(16px)
            <br />
            letter spacing: normal
          </Text>
        </div>
        <div className='flex flex-col gap-1'>
          <Text variant='xs-semibold'>Xs Semibold</Text>
          <Text as='span' className='text-muted-foreground'>
            font size: 0.75rem(12px)
            <br /> font weight: 600
            <br /> line height: 1rem(16px)
            <br />
            letter spacing: normal
          </Text>
        </div>
      </div>
    )
  },
}

export const Heading5xl: Story = {
  args: {
    children: 'Heading 5xl',
    variant: 'heading-5xl',
  },
}

export const Heading4xl: Story = {
  args: {
    children: 'Heading 4xl',
    variant: 'heading-4xl',
  },
}

export const Heading3xl: Story = {
  args: {
    children: 'Heading 3xl',
    variant: 'heading-3xl',
  },
}

export const Heading2xl: Story = {
  args: {
    children: 'Heading 2xl',
    variant: 'heading-2xl',
  },
}

export const HeadingXl: Story = {
  args: {
    children: 'Heading xl',
    variant: 'heading-xl',
  },
}

export const HeadingLg: Story = {
  args: {
    children: 'Heading lg',
    variant: 'heading-lg',
  },
}

export const HeadingMd: Story = {
  args: {
    children: 'Heading md',
    variant: 'heading-md',
  },
}

export const HeadingSm: Story = {
  args: {
    children: 'Heading sm',
    variant: 'heading-sm',
  },
}

export const HeadingXs: Story = {
  args: {
    children: 'Heading xs',
    variant: 'heading-xs',
  },
}

export const LgRegular: Story = {
  args: {
    children: 'Lg Regular',
    variant: 'lg-regular',
  },
}

export const LgSemibold: Story = {
  args: {
    children: 'Lg Semibold',
    variant: 'lg-semibold',
  },
}

export const MdRegular: Story = {
  args: {
    children: 'Md Regular',
    variant: 'md-regular',
  },
}

export const MdSemibold: Story = {
  args: {
    children: 'Md Semibold',
    variant: 'md-semibold',
  },
}

export const SmRegular: Story = {
  args: {
    children: 'Sm Regular',
    variant: 'sm-regular',
  },
}

export const SmSemibold: Story = {
  args: {
    children: 'Sm Semibold',
    variant: 'sm-semibold',
  },
}

export const XsRegular: Story = {
  args: {
    children: 'Xs Regular',
    variant: 'xs-regular',
  },
}

export const XsSemibold: Story = {
  args: {
    children: 'Xs Semibold',
    variant: 'xs-semibold',
  },
}
