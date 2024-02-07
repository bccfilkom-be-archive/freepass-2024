import type { NextFunction, Request, Response } from 'express'
import { findUserById, updateUserById, findUserByField } from '../services/user.service'
import { updateUserValidation } from '../validation/user.validation'
import type { UpdateUserForm } from '../types/user.type'
import { logger } from '../utils/logger'

export const updateUser = async (req: Request, res: Response, next: NextFunction) => {
  const id: string = res.locals.user._doc._id

  try {
    const payload: UpdateUserForm = req.body

    const { error } = updateUserValidation(payload)
    if (error) throw error

    const user = await findUserById(id)

    if (user) {
      let foundUser: boolean = false
      if (payload.nim) foundUser = !!(await findUserByField('nim', payload.nim))
      if (foundUser) throw Error('nim has been taken')

      if (payload.username) foundUser = !!(await findUserByField('username', payload.username))
      if (foundUser) throw Error('username has been taken')

      if (payload.email) foundUser = !!(await findUserByField('email', payload.email))
      if (foundUser) throw new Error('email has been taken')
    }

    await updateUserById(id, payload)
    return res.status(200).send({ status: 200, message: 'user update success', data: payload })
  } catch (error: any) {
    if (error.message.includes('has been taken')) {
      logger.error('user - update - searching in db: ', error.message)
      res.status(400).send({ status: 400, message: error.message })
    } else {
      next(error)
    }
  }
}
