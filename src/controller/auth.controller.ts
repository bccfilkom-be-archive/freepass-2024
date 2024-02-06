import type { Request, Response } from 'express'
import { createUserValidation, loginValidation } from '../validation/auth.validation'
import type { RegisterForm, LoginForm } from '../types/auth.type'
import { logger } from '../utils/logger'
import { createUser, findUserByUsername } from '../services/auth.service'
import { checkPassword, hashing } from '../utils/bcrypt'
import { signJWT } from '../utils/jwt'

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

export const login = async (req: Request, res: Response) => {
  const payload: LoginForm = req.body
  const { error, value } = loginValidation(payload)
  if (error) {
    logger.error(`auth - login - validation: ${error.details[0].message.replace(/"/g, '')}`)
    return res.status(400).send({ status: 400, message: `login failed. ${error.details[0].message.replace(/"/g, '')}` })
  }

  try {
    const payload: LoginForm = value
    const user = await findUserByUsername(payload.username)
    if (!user) {
      logger.error(`auth - login - searching in db: ${payload.username} not found`)
      return res.status(401).send({ status: 401, message: `login failed. ${payload.username} not found` })
    }

    const isValid = checkPassword(payload.password, user.password)
    if (!isValid) {
      logger.error('auth - login - searching in db: password wrong')
      return res.status(401).send({ status: 401, message: 'login failed. password is wrong' })
    }

    const token = signJWT({ ...user }, { expiresIn: '30s' })
    if (!token) {
      logger.error('auth - login - creating token')
      return res.status(500).send({ status: 500, message: 'internal server error' })
    }

    return res.status(200).send({ status: 200, message: 'login success', data: token })
  } catch (error: any) {
    logger.error(`auth - login - searching in db: ${error}`)
    return res.status(500).send({ status: 500, message: `login failed. ${error}` })
  }
}
