import { Router } from 'express'

import { updateUser } from '../../controller/user.controller'
import { requireUser } from '../../middlewares/auth'

export const userRouter: Router = Router()

userRouter.patch('/profile', requireUser, updateUser)
