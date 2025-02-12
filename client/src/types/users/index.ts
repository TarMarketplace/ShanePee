import { z } from 'zod'

const userSchema = z.object({
  id: z.number(),
  email: z.string(),
  first_name: z.string().nullable(),
  last_name: z.string().nullable(),
  gender: z.string().nullable(),
  tel: z.string().nullable(),
  address: z.object({
    house_no: z.string().nullable(),
    district: z.string().nullable(),
    province: z.string().nullable(),
    postcode: z.string().nullable(),
  }),
  payment_method: z.object({
    card_number: z.string().nullable(),
    card_owner: z.string().nullable(),
    cvv: z.string().nullable(),
    expire_date: z.string().nullable(),
  }),
})

type User = z.infer<typeof userSchema>

export { userSchema }
export type { User }
