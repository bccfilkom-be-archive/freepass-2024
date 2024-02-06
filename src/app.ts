import express, { type Application } from 'express'
import helmet from 'helmet'
import cors from 'cors'
import { successHandler, errorHandler } from './utils/morgan'
import { router } from './routes/v1/index'
import { unknownEndpoint, errorHandlerEndpoint } from './middlewares/error'

export const app: Application = express()

// set security HTTP headers
app.use(helmet())

// parse json request body
app.use(express.json())

// parse urlencoded request body
app.use(express.urlencoded({ extended: true }))

// enable cors
app.use(cors())
app.options('*', cors())

// morgan
app.use(successHandler)
app.use(errorHandler)

// v1 api routes
app.use('/v1', router)

app.use(unknownEndpoint)
app.use(errorHandlerEndpoint)
