import { type Request, type Response, type NextFunction } from 'express'
import { logger } from '../utils/logger'

export const unknownEndpoint = (request: Request, response: Response) => {
  response.status(404).send({ status: 404, message: 'unknown endpoint' })
}

export const errorHandlerEndpoint = (error: Error, request: Request, response: Response, next: NextFunction) => {
  logger.error(error.message)

  if (error.name === 'CastError') {
    return response.status(400).send({ status: 400, message: 'malformatted id' })
  } else if (error.name === 'ValidationError') {
    return response.status(400).json({ status: 400, message: error.message })
  } else if (error.name === 'JsonWebTokenError') {
    return response.status(401).json({ status: 401, message: 'invalid token' })
  } else if (error.name === 'TokenExpiredError') {
    return response.status(401).json({ status: 401, message: 'token expired' })
  } else {
    next(error)
  }
}
