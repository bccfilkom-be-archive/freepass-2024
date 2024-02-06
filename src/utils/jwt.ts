import jwt from 'jsonwebtoken'
import { config } from './config'

export const signJWT = (payload: object, options?: jwt.SignOptions | undefined) => {
  return jwt.sign(payload, config.jwtSecret, {
    ...(options && options)
  })
}
