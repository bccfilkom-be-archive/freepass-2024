import { Post } from '../models/post.model'
import type { CreatePostForm, UpdatePostForm } from '../types/post.type'

export const createPostForId = async (payload: CreatePostForm, candidateId: string) => {
  return await Post.create({ ...payload, candidateId })
}

export const findPostById = async (id: string) => {
  return await Post.findById(id)
}

export const findPostByField = async (field: string, value: string) => {
  return await Post.findOne({ [field]: value })
}

export const updatePostById = async (id: string, payload: UpdatePostForm) => {
  return await Post.findByIdAndUpdate(id, payload, { new: true })
}

export const deletePostById = async (id: string) => {
  return await Post.findByIdAndDelete(id)
}

export const getAllPosts = async () => {
  return await Post.find({})
}
