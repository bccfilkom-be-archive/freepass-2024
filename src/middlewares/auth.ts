import type { Request, Response, NextFunction } from 'express'

export const requireUser = (req: Request, res: Response, next: NextFunction) => {
  const user = res.locals.user
  if (!user) {
    return res.status(403).send({ status: 403, message: 'forbidden' })
  }
  next()
}
