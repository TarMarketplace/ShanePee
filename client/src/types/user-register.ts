import { z } from 'zod'

const userDetailSchema = z.object({
  name: z.string(),
  lastname: z.string(),
  email: z.string(),
  phone: z.string(),
  gender: z.string(),
  username: z.string(),
  password: z.string(),
  passwordConfirm: z.string(),
})

type UserRegisterDetail = z.infer<typeof userDetailSchema>
type ErrorRegister = {
  [key in keyof UserRegisterDetail]: string
}

export { userDetailSchema }
export type { UserRegisterDetail, ErrorRegister }
