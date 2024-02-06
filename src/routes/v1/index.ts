import express, { type Router } from 'express'
import { docsRouter } from './docs.route'
import { authRouter } from './auth.route'

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
  }
]

defaultRoutes.forEach((route: Route) => {
  router.use(route.path, route.route)
})
