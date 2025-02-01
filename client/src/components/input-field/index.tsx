'use client'

import { Input } from '@material-tailwind/react'

const InputField = (props) => {
  return (
    <div className='h-14'>
      <div className='bg-white'>
        <Input
          type={props.type ? props.type : 'text'}
          size={props.size}
          label={props.label}
          name={props.name}
          value={props.value}
          onChange={props.onChange}
          error={props.error}
        />
      </div>
      {props.error && <p className='text-sm text-red-500'>{props.error}</p>}
    </div>
  )
}

export { InputField }
