import express, { type Router } from 'express'
import { docsRouter } from './docs.route'

export const router: Router = express.Router()

interface Route {
  path: string
  route: Router
}

const defaultRoutes: Route[] = [
  {
    path: '/docs',
    route: docsRouter
  }
]

defaultRoutes.forEach((route: Route) => {
  router.use(route.path, route.route)
})
