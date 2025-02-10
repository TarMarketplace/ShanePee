import { z } from 'zod'

const userSchema = z.object({
  id: z.number(),
  first_name: z.string(),
  last_name: z.string(),
  gender: z.string(),
  email: z.string(),
  tel: z.string(),
  address: z.object({
    house_no: z.string(),
    district: z.string(),
    province: z.string(),
    postcode: z.string(),
  }),
  payment_method: z.object({
    card_number: z.string(),
    card_owner: z.string(),
    cvv: z.string(),
    expire_date: z.string(),
  }),
})

type User = z.infer<typeof userSchema>

export { userSchema }
export type { User }
