import type { CreateCommentForm } from '../types/comment.type'
import { Comment } from '../models/comment.model'

export const createCommentForId = async (payload: CreateCommentForm, postId: string, userId: string) => {
  return await Comment.create({ ...payload, postId, userId })
}

export const findCommentById = async (id: string) => {
  return await Comment.findById(id)
}

export const deleteCommentById = async (id: string) => {
  return await Comment.findByIdAndDelete(id)
}
