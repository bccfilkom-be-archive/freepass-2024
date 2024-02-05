import { type Request, type Response, type NextFunction } from 'express'
import { logger } from '../utils/logger'

export const unknownEndpoint = (request: Request, response: Response): void => {
  response.status(404).send({ error: 'unknown endpoint' })
}

export const errorHandler = (
  error: Error,
  request: Request,
  response: Response,
  next: NextFunction
): Response<any, Record<string, any>> | undefined => {
  logger.error(error.message)

  if (error.name === 'CastError') {
    return response.status(400).send({ error: 'malformatted id' })
  } else if (error.name === 'ValidationError') {
    return response.status(400).json({ error: error.message })
  } else if (error.name === 'JsonWebTokenError') {
    return response.status(401).json({ error: 'invalid token' })
  } else if (error.name === 'TokenExpiredError') {
    return response.status(401).json({ error: 'token expired' })
  } else {
    next(error)
  }
}
