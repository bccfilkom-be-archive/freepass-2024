import { Router } from 'express'
import { promoteUser } from '../../controller/admin.controller'
import { requireAdmin } from '../../middlewares/auth'

export const adminRouter: Router = Router()

adminRouter.post('/:id', requireAdmin, promoteUser)
