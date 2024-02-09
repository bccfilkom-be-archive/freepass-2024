import type { CreateCommentForm } from '../types/comment.type'
import { Comment } from '../models/comment.model'
import { findUserById } from './user.service'
import { findPostById } from './post.service'

export const createCommentForId = async (payload: CreateCommentForm, postId: string, userId: string) => {
  return await Comment.create({ ...payload, postId, userId })
}

export const findCommentById = async (id: string) => {
  return await Comment.findById(id)
}

export const deleteCommentById = async (id: string) => {
  const comment = await findCommentById(id)
  if (!comment) return comment

  const userId = comment.userId._id.toString()
  const user = await findUserById(userId)
  if (user) {
    user.comments = user.comments.filter((c) => c._id.toString() !== comment._id.toString())
    user.commentedPosts = user.commentedPosts.filter((post) => post._id.toString() !== comment.postId.toString())
    await user.save()
  }

  const postId = comment.postId._id.toString()
  const post = await findPostById(postId)
  if (post) {
    post.comments = post.comments.filter((c) => c._id.toString() !== comment._id.toString())
    await post.save()
  }

  return await Comment.findByIdAndDelete(id)
}
