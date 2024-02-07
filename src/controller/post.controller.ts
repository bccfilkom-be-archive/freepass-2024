import type { Request, Response, NextFunction } from 'express'
import { createFormValidation, updatePostValidation } from '../validation/post.validation'
import type { CreatePostForm, UpdatePostForm } from '../types/post.type'
import { createPostForId, deletePostById, findPostById, updatePostById } from '../services/post.service'
import { logger } from '../utils/logger'
import { findCandidateByField } from '../services/candidate.service'

export const createPost = async (req: Request, res: Response, next: NextFunction) => {
  try {
    const payload: CreatePostForm = req.body
    const { error } = createFormValidation(payload)
    if (error) throw error

    const userId: string = res.locals.user._doc._id
    const candidate = await findCandidateByField('userId', userId)
    if (candidate) {
      const post = await createPostForId(payload, candidate._id.toString())
      candidate.posts = candidate.posts.concat(post._id)
      await candidate.save()
    }

    return res.status(201).send({ status: 201, message: 'create post success' })
  } catch (error: any) {
    if (error.message.includes('has been taken')) {
      logger.error('post - create - searching in db: ', error.message)
      res.status(400).send({ status: 400, message: error.message })
    } else {
      next(error)
    }
  }
}

export const updatePost = async (req: Request, res: Response, next: NextFunction) => {
  const postId = req.params.postId

  try {
    const payload: UpdatePostForm = req.body
    const { error } = updatePostValidation(payload)
    if (error) throw error

    const post = await findPostById(postId)
    if (!post) throw new Error('post not found')

    await updatePostById(postId, payload)
    return res.status(200).send({ status: 200, message: 'post update success', data: payload })
  } catch (error: any) {
    if (error.message.includes('not found')) {
      logger.error('post - update - searching in db: ', error.message)
      res.status(404).send({ status: 404, message: error.message })
    } else {
      next(error)
    }
  }
}

export const deletePost = async (req: Request, res: Response, next: NextFunction) => {
  const postId = req.params.postId

  try {
    const post = await findPostById(postId)
    if (!post) throw new Error('post not found')

    await deletePostById(postId)

    const userId: string = res.locals.user._doc._id
    const candidate = await findCandidateByField('userId', userId)
    if (candidate) {
      candidate.posts = candidate.posts.filter((post) => post._id.toString() !== postId.toString())
      await candidate.save()
    }
    return res.status(200).send({ status: 200, message: 'delete post success' })
  } catch (error: any) {
    if (error.message.includes('not found')) {
      logger.error('post - delete - searching in db: ', error.message)
      res.status(404).send({ status: 404, message: error.message })
    } else {
      next(error)
    }
  }
}

export const viewPost = async (req: Request, res: Response, next: NextFunction) => {
  const postId = req.params.id

  try {
    const post = await findPostById(postId)
    if (!post) throw new Error('post not found')

    return res.status(200).send({ status: 200, message: 'view post success', data: post })
  } catch (error: any) {
    if (error.message.includes('not found')) {
      logger.error('post - delete - searching in db: ', error.message)
      res.status(404).send({ status: 404, message: error.message })
    } else {
      next(error)
    }
  }
}
