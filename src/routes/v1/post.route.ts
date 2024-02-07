import { Router } from 'express'
import { requireCandidate } from '../../middlewares/auth'
import { createPost, updatePost, deletePost } from '../../controller/post.controller'

export const postRouter: Router = Router()

postRouter.post('/', requireCandidate, createPost)
postRouter.patch('/:postId', requireCandidate, updatePost)
postRouter.delete('/:postId', requireCandidate, deletePost)
