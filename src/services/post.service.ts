import { Post } from '../models/post.model'
import type { CreatePostForm } from '../types/post.type'

export const createPostForId = async (payload: CreatePostForm, candidateId: string) => {
  return await Post.create({ ...payload, candidateId })
}

export const findPostById = async (id: string) => {
  return await Post.findById(id)
}

export const findPostByField = async (field: string, value: string) => {
  return await Post.findOne({ [field]: value })
}
