import express, { type Application, type NextFunction, type Request, type Response, type RequestHandler } from 'express'
import helmet from 'helmet'
import cors from 'cors'
import { type ParsedQs } from 'qs'
import { router } from './routes/v1/index'
import { logger } from './utils/logger'
import { unknownEndpoint, errorHandler } from './middlewares/error'

const app: Application = express()
const hostname: string = '0.0.0.0'
const port: number = 4000

// set security HTTP headers
app.use(helmet())

// parse json request body
app.use(express.json())

// parse urlencoded request body
app.use(express.urlencoded({ extended: true }))

// enable cors
app.use(cors())
const corsHandler: RequestHandler<Record<any, any>, any, any, ParsedQs, Record<string, any>> = cors()
app.options('*', corsHandler)

// v1 api routes
app.use('/v1', router)

app.use('/', (req: Request, res: Response, next: NextFunction) => {
  res.json({
    data: 'Menyala abangkuh🔥🔥🔥'
  })
})

app.use(unknownEndpoint)
app.use(errorHandler)

app.listen(port, hostname, () => {
  logger.info(`Server is listening on port ${port}`)
})
