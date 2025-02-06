import { z } from 'zod'

const userSchema = z.object({
  id: z.string(),
  name: z.string(),
})

type User = z.infer<typeof userSchema>

export { userSchema }
export type { User }
