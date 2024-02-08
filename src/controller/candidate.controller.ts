import type { NextFunction, Request, Response } from 'express'
import { findCandidateById, getAllCandidates } from '../services/candidate.service'

export const viewAllCandidates = async (req: Request, res: Response, next: NextFunction) => {
  const candidates = await getAllCandidates()
  return res.status(200).send({
    status: 200,
    message: 'get all candidates success',
    data: { candidates: [...candidates], length: candidates.length }
  })
}

export const viewCandidate = async (req: Request, res: Response, next: NextFunction) => {
  const candidateId = req.params.id

  try {
    const id = await findCandidateById(candidateId)
    if (!id) throw new Error('id not found')

    return res.status(200).send({ status: 200, message: 'view candidate by id success', data: id })
  } catch (error: any) {
    if (error.message.includes('not found')) {
      res.status(404).send({ status: 404, message: error.message })
    } else {
      next(error)
    }
  }
}

export const viewCandidatePosts = async (req: Request, res: Response, next: NextFunction) => {
  const candidateId = req.params.id

  try {
    const id = await findCandidateById(candidateId)
    if (!id) throw new Error('id not found')

    const candidate = await findCandidateById(candidateId)
    if (candidate) {
      const posts = candidate.posts
      return res
        .status(200)
        .send({ status: 200, message: "view candidate's posts success", data: { posts, length: posts.length } })
    }
  } catch (error: any) {
    if (error.message.includes('not found')) {
      res.status(404).send({ status: 404, message: error.message })
    } else {
      next(error)
    }
  }
}
