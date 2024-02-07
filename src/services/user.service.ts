import type { UpdateUserForm } from '../types/user.type'
import { User } from '../models/user.model'

export const findUserById = async (id: string) => {
  return await User.findById(id)
}

export const findUserByField = async (field: string, value: string) => {
  return await User.findOne({ [field]: value })
}

export const updateUserById = async (id: string, payload: UpdateUserForm) => {
  return await User.findByIdAndUpdate(id, payload, { new: true })
}

export const findUserByIdAndPromote = async (id: string) => {
  return await User.findByIdAndUpdate(id, { role: 'candidate' }, { new: true })
}
