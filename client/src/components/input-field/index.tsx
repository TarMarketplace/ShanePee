'use client'

import { Input } from '@material-tailwind/react'

interface InputFieldProps {
  type?: string
  size?: 'sm' | 'md' | 'lg'
  label?: string
  name?: string
  value?: string | number
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
  error?: string
}

const InputField = ({
  type = 'text',
  size,
  label,
  name,
  value,
  onChange,
  error,
}: InputFieldProps) => {
  return (
    <div className='h-14'>
      <div className='bg-white'>
        <Input
          type={type}
          size={size}
          label={label}
          name={name}
          value={value}
          onChange={onChange}
          error={!!error}
        />
      </div>
      {error && typeof error === 'string' && (
        <p className='text-sm text-red-500'>{error}</p>
      )}
    </div>
  )
}

export { InputField }
