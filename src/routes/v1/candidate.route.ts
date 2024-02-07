import { Router } from 'express'

import { viewAllCandidates, viewCandidate } from '../../controller/candidate.controller'
import { requireAdmin } from '../../middlewares/auth'

export const candidateRouter: Router = Router()

candidateRouter.get('/', requireAdmin, viewAllCandidates)
candidateRouter.get('/:id', requireAdmin, viewCandidate)
