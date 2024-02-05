import express, { type Application, type NextFunction, type Request, type Response } from 'express'

const app: Application = express()
const port: number = 4000

app.use('/', (req: Request, res: Response, next: NextFunction) => {
  res.json({
    data: 'Menyala abangkuh🔥🔥🔥'
  })
})

app.listen(port, () => {
  console.log(`Server is listening on port ${port}`)
})
