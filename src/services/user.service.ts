import type { UpdateUserForm } from '../types/user.type'
import { User } from '../models/user.model'
import { deleteCommentById } from './comment.service'

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

export const getAllUsers = async () => {
  return await User.find({})
}

export const deleteUserById = async (id: string) => {
  const user = await findUserById(id)
  if (!user) return user

  const comments = user.comments
  const deleteCommentPromises = comments.map(async (comment) => {
    await deleteCommentById(comment._id.toString())
  })
  await Promise.all(deleteCommentPromises)

  return await User.findByIdAndDelete(id)
}
