import type { Request, Response, NextFunction } from 'express'
import type { CreateCommentForm } from '../types/comment.type'
import { createFormValidation } from '../validation/comment.validation'
import { findPostById } from '../services/post.service'
import { createCommentForId } from '../services/comment.service'
import { findUserById } from '../services/user.service'

export const createComment = async (req: Request, res: Response, next: NextFunction) => {
  const postId = req.params.postId
  const userId: string = res.locals.user._doc._id

  try {
    const post = await findPostById(postId)
    if (!post) throw new Error('post not found')

    const payload: CreateCommentForm = req.body
    const { error } = createFormValidation(payload)
    if (error) throw error

    const comment = await createCommentForId(payload, postId, userId)
    const user = await findUserById(userId)
    if (user) {
      user.commentedPosts = user.commentedPosts.concat(post._id)
      user.comments = user.comments.concat(comment._id)
      await user.save()

      post.comments = post.comments.concat(comment._id)
      await post.save()
    }
    return res.status(201).send({ status: 201, message: 'create comment success' })
  } catch (error: any) {
    if (error.message.includes('not found')) {
      res.status(400).send({ status: 400, message: error.message })
    } else {
      next(error)
    }
  }
}
