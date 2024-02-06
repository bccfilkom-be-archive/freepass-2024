import type { Request, Response } from 'express'
import { createUserValidation } from '../validation/auth.validation'
import type { UserType } from '../types/user.type'
import { logger } from '../utils/logger'
import { createUser } from '../services/auth.service'
import { hashing } from '../utils/bcrypt'

export const register = async (req: Request, res: Response) => {
  const payload: UserType = req.body
  logger.info('auth - register: ', payload)
  const { error, value } = createUserValidation(payload)
  if (error) {
    const errorMessage = error.details.map((details) => details.message).join(', ')
    logger.error('auth - register - validation: ', errorMessage)
    return res.status(500).send({ status: 500, message: 'create user failed', data: error })
  }

  try {
    const payload: UserType = value
    payload.password = hashing(payload.password)
    const user = await createUser(payload)
    return res.status(201).send({ status: 201, message: 'create user success', data: user })
  } catch (error) {
    logger.error('auth - register - saving to db: ', error)
    return res.status(500).send({ status: 500, message: 'create user failed', data: error })
  }
}
