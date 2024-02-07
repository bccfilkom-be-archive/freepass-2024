import { Router } from 'express'
import { requireCandidate } from '../../middlewares/auth'
import { createPost, updatePost, deletePost, viewPost } from '../../controller/post.controller'

export const postRouter: Router = Router()

postRouter.get('/:id', viewPost)
postRouter.post('/', requireCandidate, createPost)
postRouter.patch('/:postId', requireCandidate, updatePost)
postRouter.delete('/:postId', requireCandidate, deletePost)
