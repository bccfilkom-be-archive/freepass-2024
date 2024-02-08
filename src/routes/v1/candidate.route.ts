import { Router } from 'express'

import { viewAllCandidates, viewCandidate, viewCandidatePosts } from '../../controller/candidate.controller'
import { requireAdmin } from '../../middlewares/auth'

export const candidateRouter: Router = Router()

candidateRouter.get('/', requireAdmin, viewAllCandidates)
candidateRouter.get('/:id', requireAdmin, viewCandidate)
candidateRouter.get('/:id/post', requireAdmin, viewCandidatePosts)
