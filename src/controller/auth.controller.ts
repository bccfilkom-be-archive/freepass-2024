import type { Request, Response } from 'express'
import { createUserValidation } from '../validation/auth.validation'
import type { RegisterForm } from '../types/auth.type'
import { logger } from '../utils/logger'
import { createUser } from '../services/auth.service'
import { hashing } from '../utils/bcrypt'

export const register = async (req: Request, res: Response) => {
  const payload: RegisterForm = req.body
  const { error, value } = createUserValidation(payload)
  if (error) {
    logger.error(`auth - register - validation: ${error.details[0].message.replace(/"/g, '')}`)
    return res
      .status(400)
      .send({ status: 400, message: `create user failed. ${error.details[0].message.replace(/"/g, '')}` })
  }

  try {
    const payload: RegisterForm = value
    payload.password = hashing(payload.password)
    const user = await createUser(payload)
    return res.status(201).send({ status: 201, message: `create user success. please login ${user.username}` })
  } catch (error: any) {
    const taken: object = error.keyValue
    logger.error(`auth - register - saving to db: ${Object.keys(taken)[0]} has been taken`)
    return res.status(400).send({ status: 400, message: `create user failed. ${Object.keys(taken)[0]} has been taken` })
  }
}
