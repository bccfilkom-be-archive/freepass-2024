import { Router } from 'express'
import { promoteUser, deleteUser, deleteCandidate, deletePost, deleteComment } from '../../controller/admin.controller'
import { requireAdmin } from '../../middlewares/auth'

export const adminRouter: Router = Router()

adminRouter.post('/:id', requireAdmin, promoteUser)
adminRouter.delete('/user/:id', requireAdmin, deleteUser)
adminRouter.delete('/candidate/:id', requireAdmin, deleteCandidate)
adminRouter.delete('/post/:id', requireAdmin, deletePost)
adminRouter.delete('/comment/:id', requireAdmin, deleteComment)
