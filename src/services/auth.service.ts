import type { RegisterForm } from '../types/auth.type'
import { User } from '../models/user.model'

export const createUser = async (payload: RegisterForm) => {
  return await User.create(payload)
}
