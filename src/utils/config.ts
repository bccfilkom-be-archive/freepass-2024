import dotenv from 'dotenv'
dotenv.config()

const PORT: number = Number(process.env.PORT) ?? 4000
const MONGODB_URI = process.env.MONGODB_URI ?? 'MongoDB URI not found'
const JWT_SECRET = process.env.JWT_SECRET ?? 'thisisverysecret'
const JWT_ACCESS_EXPIRATION = process.env.JWT_ACCESS_EXPIRATION
const NODE_ENV = process.env.NODE_ENV ?? 'development'

export const config = {
  port: PORT,
  mongodbUri: MONGODB_URI,
  jwtSecret: JWT_SECRET,
  jwtAccessExpiration: JWT_ACCESS_EXPIRATION,
  nodeEnv: NODE_ENV
}
