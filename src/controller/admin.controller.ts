import type { NextFunction, Request, Response } from 'express'
import { deleteUserById, findUserById, findUserByIdAndPromote } from '../services/user.service'
import { createCandidate, deleteCandidateById, findCandidateByField } from '../services/candidate.service'
import { deleteCommentById } from '../services/comment.service'
import { deletePostById } from '../services/post.service'
import type { CreateElectionForm } from '../types/election.type'
import { createElectionValidation } from '../validation/election.validation'
import { createElectionService, findElectionById, getAllElections } from '../services/election.service'
import { stringtoDate } from '../utils/date'

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

export const createElection = async (req: Request, res: Response, next: NextFunction) => {
  try {
    const startDateString: string = req.body.startDate
    const endDateString: string = req.body.endDate

    const startDate = stringtoDate(startDateString)
    const endDate = stringtoDate(endDateString)

    const payload: CreateElectionForm = {
      startDate,
      endDate
    }

    const { error } = createElectionValidation(payload)
    if (error) throw error

    await createElectionService(payload)
    return res.status(201).send({ status: 201, message: 'create election success' })
  } catch (error: any) {
    next(error)
  }
}

export const viewAllElections = async (req: Request, res: Response, next: NextFunction) => {
  const elections = await getAllElections()
  return res.status(200).send({
    status: 200,
    message: 'get all elections success',
    data: { elections: [...elections], length: elections.length }
  })
}

export const viewElection = async (req: Request, res: Response, next: NextFunction) => {
  const id = req.params.id

  try {
    const election = await findElectionById(id)
    if (!election) throw new Error('election not found')

    return res.status(200).send({ status: 200, message: 'view election success', data: election })
  } catch (error: any) {
    if (error.message.includes('not found')) {
      res.status(404).send({ status: 404, message: error.message })
    } else {
      next(error)
    }
  }
}
