import { z } from 'zod'

const productSchema = z.object({
  id: z.string(),
  name: z.string(),
  price: z.number(),
  discount: z.number().nullable(),
  rating: z.number(),
  location: z.string(),
  image: z.string(),
})

type Product = z.infer<typeof productSchema>

export { productSchema }
export type { Product }
