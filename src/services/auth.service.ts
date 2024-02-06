import type { UserType } from '../types/user.type'
import { User } from '../models/user.model'

export const createUser = async (payload: UserType) => {
  return await User.create(payload)
}
