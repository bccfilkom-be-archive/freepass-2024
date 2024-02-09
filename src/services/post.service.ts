import { Post } from '../models/post.model'
import type { CreatePostForm, UpdatePostForm } from '../types/post.type'
import { findCandidateById } from './candidate.service'
import { deleteCommentById } from './comment.service'

export const createPostForId = async (payload: CreatePostForm, candidateId: string) => {
  return await Post.create({ ...payload, candidateId })
}

export const findPostById = async (id: string) => {
  return await Post.findById(id).populate('comments')
}

export const findPostByField = async (field: string, value: string) => {
  return await Post.findOne({ [field]: value })
}

export const updatePostById = async (id: string, payload: UpdatePostForm) => {
  return await Post.findByIdAndUpdate(id, payload, { new: true })
}

export const deletePostById = async (id: string) => {
  const post = await findPostById(id)
  if (!post) return post

  const candidate = await findCandidateById(post.candidateId.toString())
  if (candidate) {
    candidate.posts = candidate.posts.filter((p) => p._id.toString() !== post._id.toString())
    await candidate.save()
  }

  const comments = post.comments
  const deleteCommentPromises = comments.map(async (comment) => {
    await deleteCommentById(comment._id.toString())
  })
  await Promise.all(deleteCommentPromises)

  return await Post.findByIdAndDelete(id)
}

export const getAllPosts = async () => {
  return await Post.find({})
}
