import type { NextFunction, Request, Response } from 'express'
import { deleteUserById, findUserById, findUserByIdAndPromote } from '../services/user.service'
import { createCandidate, deleteCandidateById, findCandidateByField } from '../services/candidate.service'
import { deleteCommentById } from '../services/comment.service'
import { deletePostById } from '../services/post.service'

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
    const candidate = await findCandidateByField('userId', id)
    if (candidate) await deleteCandidateById(candidate._id.toString())

    await deleteUserById(id)

    return res.status(200).send({ status: 200, message: 'delete user success' })
  } catch (error) {
    next(error)
  }
}

export const deleteCandidate = async (req: Request, res: Response, next: NextFunction) => {
  const id = req.params.id

  try {
    await deleteCandidateById(id)

    return res.status(200).send({ status: 200, message: 'delete candidate success' })
  } catch (error) {
    next(error)
  }
}

export const deletePost = async (req: Request, res: Response, next: NextFunction) => {
  const id = req.params.id

  try {
    await deletePostById(id)

    return res.status(200).send({ status: 200, message: 'delete post success' })
  } catch (error) {
    next(error)
  }
}

export const deleteComment = async (req: Request, res: Response, next: NextFunction) => {
  const id = req.params.id

  try {
    await deleteCommentById(id)

    return res.status(200).send({ status: 200, message: 'delete comment success' })
  } catch (error) {
    next(error)
  }
}
