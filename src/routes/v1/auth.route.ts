import { Router } from 'express'

import { register } from '../../controller/auth.controller'

export const authRouter: Router = Router()

authRouter.post('/register', register)
