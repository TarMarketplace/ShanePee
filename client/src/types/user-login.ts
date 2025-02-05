import { z } from 'zod'

const userLoginSchema = z.object({
  username: z.string(),
  password: z.string(),
})

type UserLoginDetail = z.infer<typeof userLoginSchema>
type ErrorLogin = {
  [key in keyof UserLoginDetail]: string
}

export { userLoginSchema }
export type { UserLoginDetail, ErrorLogin }
