import type { Request, Response, NextFunction } from 'express'
import { createFormValidation } from '../validation/post.validation'
import type { CreatePostForm } from '../types/post.type'
import { createPostForId } from '../services/post.service'
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
      logger.error('user - update - searching in db: ', error.message)
      res.status(400).send({ status: 400, message: error.message })
    } else {
      next(error)
    }
  }
}
