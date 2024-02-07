import express, { type Router } from 'express'
import { docsRouter } from './docs.route'
import { authRouter } from './auth.route'
import { userRouter } from './user.route'

export const router: Router = express.Router()

interface Route {
  path: string
  route: Router
}

const defaultRoutes: Route[] = [
  {
    path: '/docs',
    route: docsRouter
  },
  {
    path: '/auth',
    route: authRouter
  },
  {
    path: '/user',
    route: userRouter
  }
]

defaultRoutes.forEach((route: Route) => {
  router.use(route.path, route.route)
})