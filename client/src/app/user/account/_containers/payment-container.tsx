import { zodResolver } from '@hookform/resolvers/zod'
import { Icon } from '@iconify/react/dist/iconify.js'
import { type SubmitHandler, useForm } from 'react-hook-form'
import { z } from 'zod'

import { Text } from '@/components/text'

import { PaymentForm } from '../_components/payment-form'

const paymentFormSchema = z.object({
  cardNumber: z
    .string()
    .length(16, 'Card number must be exactly 16 digits')
    .regex(/^\d{16}$/, 'Card number must contain only digits'),
  cardHolderName: z.string().min(1, 'Cardholder name is required'),
  expirationDate: z
    .string()
    .length(5, 'Expiration date must be in MM/YY format')
    .regex(/^(0[1-9]|1[0-2])\/\d{2}$/, 'Invalid expiration date format'),
  cvv: z
    .string()
    .length(3, 'CVV must be exactly 3 digits')
    .regex(/^\d{3}$/, 'CVV must contain only digits'),
})

export type PaymentFormSchema = z.infer<typeof paymentFormSchema>

export function PaymentContainer() {
  const form = useForm<PaymentFormSchema>({
    resolver: zodResolver(paymentFormSchema),
    defaultValues: {
      cardNumber: '',
      cardHolderName: '',
      expirationDate: '',
      cvv: '',
    },
  })

  const onSubmit: SubmitHandler<PaymentFormSchema> = (data) => {
    console.log(data)
  }

  return (
    <div className='m-6 flex w-full flex-col gap-8'>
      <div className='flex items-center gap-2'>
        <Icon icon='famicons:card' className='size-10' />
        <Text variant='heading-lg'>ช่องทางการชำระเงิน</Text>
      </div>
      <PaymentForm onSubmit={onSubmit} form={form} />
    </div>
  )
}
