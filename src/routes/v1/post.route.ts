import { Router } from 'express'
import { requireCandidate } from '../../middlewares/auth'
import { createPost } from '../../controller/post.controller'

export const postRouter: Router = Router()

postRouter.post('/', requireCandidate, createPost)
