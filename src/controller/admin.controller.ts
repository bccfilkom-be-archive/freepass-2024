import type { NextFunction, Request, Response } from 'express'
import { findUserById, findUserByIdAndPromote } from '../services/user.service'

export const promoteUser = async (req: Request, res: Response, next: NextFunction) => {
  const id: string = req.params.id

  try {
    const user = await findUserById(id)

    if (user) {
      if (user.role === 'admin') throw new Error('cannot promote admin')
      await findUserByIdAndPromote(id)
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
