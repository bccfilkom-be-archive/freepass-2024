import { Router } from 'express'

import { updateUser, viewAllUsers } from '../../controller/user.controller'
import { requireAdmin, requireUser } from '../../middlewares/auth'

export const userRouter: Router = Router()

userRouter.get('/', requireAdmin, viewAllUsers)
userRouter.patch('/profile', requireUser, updateUser)
