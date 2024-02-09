import type { NextFunction, Request, Response } from 'express'
import { deleteUserById, findUserById, findUserByIdAndPromote } from '../services/user.service'
import {
  createCandidate,
  deleteCandidateById,
  findCandidateByField,
  findCandidateById
} from '../services/candidate.service'
import { deleteCommentById, findCommentById } from '../services/comment.service'
import { deletePostById, findPostById } from '../services/post.service'

export const promoteUser = async (req: Request, res: Response, next: NextFunction) => {
  const id: string = req.params.id

  try {
    const user = await findUserById(id)

    if (user) {
      if (user.role === 'admin') throw new Error('cannot promote admin')
      await findUserByIdAndPromote(id)
      await createCandidate(id)
    } else {
      throw new Error('cannot promote user')
    }
    return res.status(200).send({ status: 200, message: 'user promote success' })
  } catch (error: any) {
    if (error.message.includes('cannot promote')) {
      res.status(400).send({ status: 400, message: error.message })
    } else {
      next(error)
    }
  }
}

export const deleteUser = async (req: Request, res: Response, next: NextFunction) => {
  const id = req.params.id

  try {
    const user = await findUserById(id)

    if (user) {
      if (user.role === 'user') {
        const commentsIds = user.comments
        commentsIds.forEach(async (commentId) => {
          const comment = await findCommentById(commentId.toString())
          if (comment) {
            const postId = comment.postId
            const post = await findPostById(postId._id.toString())
            if (post) {
              post.comments = post.comments.filter((comment) => comment._id.toString() !== commentId._id.toString())
              await post.save()
              await deleteCommentById(commentId.toString())
            }
          }
        })
        await deleteUserById(id)
      } else if (user.role === 'candidate') {
        const candidate = await findCandidateByField('userId', id)
        if (candidate) {
          const postIds = candidate.posts
          postIds.forEach(async (postId) => {
            const post = await findPostById(postId.toString())
            if (post) {
              const commentIds = post.comments
              commentIds.forEach(async (commentId) => {
                const comment = await findCommentById(commentId._id.toString())
                if (comment) {
                  const userId = comment.userId
                  const user = await findUserById(userId._id.toString())
                  if (user) {
                    user.comments = user.comments.filter(
                      (comment) => comment._id.toString() !== commentId._id.toString()
                    )
                    user.commentedPosts = user.commentedPosts.filter(
                      (post) => post._id.toString() !== comment.postId.toString()
                    )
                    await user.save()
                  }

                  await deleteCommentById(comment._id.toString())
                }
              })

              await deletePostById(post._id.toString())
            }
          })
          await deleteCandidateById(candidate._id.toString())
        }

        await deleteUserById(id)
      }
    }
    return res.status(200).send({ status: 200, message: 'delete user success' })
  } catch (error) {
    next(error)
  }
}

export const deleteCandidate = async (req: Request, res: Response, next: NextFunction) => {
  const id = req.params.id

  try {
    const candidate = await findCandidateById(id)

    if (candidate) {
      const postIds = candidate.posts
      postIds.forEach(async (postId) => {
        const post = await findPostById(postId._id.toString())
        if (post) {
          const commentIds = post.comments
          commentIds.forEach(async (commentId) => {
            const comment = await findCommentById(commentId._id.toString())
            if (comment) {
              const userId = comment.userId
              const user = await findUserById(userId._id.toString())
              if (user) {
                user.comments = user.comments.filter((comment) => comment._id.toString() !== commentId._id.toString())
                user.commentedPosts = user.commentedPosts.filter(
                  (post) => post._id.toString() !== comment.postId.toString()
                )
                await user.save()
              }

              await deleteCommentById(comment._id.toString())
            }
          })

          await deletePostById(post._id.toString())
        }
      })
      await deleteCandidateById(candidate._id.toString())
      await deleteUserById(candidate.userId.toString())
    }
    return res.status(200).send({ status: 200, message: 'delete candidate success' })
  } catch (error) {
    next(error)
  }
}
