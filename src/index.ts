import express, { type Application, type NextFunction, type Request, type Response } from 'express'
import { router } from './routes/v1/index'

const app: Application = express()
const port: number = 4000

app.use('/v1', router)

app.use('/', (req: Request, res: Response, next: NextFunction) => {
  res.json({
    data: 'Menyala abangkuh🔥🔥🔥'
  })
})

app.listen(port, () => {
  console.log(`Server is listening on port ${port}`)
})
