import supertest from 'supertest'
import { MongoMemoryServer } from 'mongodb-memory-server'
import mongoose from 'mongoose'
import { app } from '../app'
import { User } from '../models/user.model'
import type { RegisterForm } from '../types/auth.type'
import { hashing } from '../utils/bcrypt'
import { findUserByField } from '../services/user.service'
import { findPostByField } from '../services/post.service'
import { findCandidateByField } from '../services/candidate.service'
import type { CreateCommentForm } from '../types/comment.type'

describe('commentRoutes', () => {
  beforeAll(async () => {
    const mongoServer = await MongoMemoryServer.create()
    await mongoose.connect(mongoServer.getUri())
  })

  afterAll(async () => {
    await mongoose.disconnect()
    await mongoose.connection.close()
  })

  describe('post /v1/post/:id', () => {
    let newCandidate: RegisterForm
    let newUser: RegisterForm
    let token: string
    let postId: string
    beforeAll(async () => {
      newCandidate = {
        fullName: 'valid full name',
        username: 'validusername',
        nim: '231502001110111',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newCandidate).expect(201)

      newUser = {
        fullName: 'valid full name',
        username: 'validusername2',
        nim: '2315020011101112',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail2@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newUser).expect(201)

      const password = hashing('password')
      const admin = new User({
        fullName: 'admin',
        nim: '0000001',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'admin@gmail.com',
        username: 'admin',
        password,
        role: 'admin'
      })
      await admin.save()

      token = (
        await supertest(app).post('/v1/auth/login').send({ username: 'admin', password: 'password' }).expect(200)
      ).body.data

      const user = await findUserByField('username', newCandidate.username)
      await supertest(app).post(`/v1/admin/${user?._id}`).set('Authorization', `Bearer ${token}`).expect(200)
      token = (
        await supertest(app)
          .post('/v1/auth/login')
          .send({ username: newCandidate.username, password: newCandidate.password })
          .expect(200)
      ).body.data
      await supertest(app).post('/v1/post/').set('Authorization', `Bearer ${token}`).send(payload)
      if (user) {
        const candidate = await findCandidateByField('userId', user._id.toString())
        if (candidate) {
          const post = await findPostByField('candidateId', candidate._id.toString())
          if (post) postId = post._id.toString()
        }
      }
    })

    let payload: CreateCommentForm
    beforeEach(() => {
      payload = {
        caption: 'valid caption'
      }
    })

    test("should return 201 if request data is ok and logged user's role is correct", async () => {
      token = (
        await supertest(app)
          .post('/v1/auth/login')
          .send({ username: newUser.username, password: newUser.password })
          .expect(200)
      ).body.data
      const res = await supertest(app).post(`/v1/post/:${postId}`).set('Authorization', `Bearer ${token}`).send(payload)

      let user = await findUserByField('username', newCandidate.username)
      if (user) {
        const candidate = await findCandidateByField('userId', user._id.toString())
        if (candidate) {
          const post = await findPostByField('candidateId', candidate._id.toString())
          if (post) {
            user = await findUserByField('username', newUser.username)
            if (user) {
              expect(res.body.status).toBe(200)
              expect(post.comments).toStrictEqual([post._id.toString()])
              expect(user.comments).toStrictEqual([post._id.toString()])
            }
          }
        }
      }
    })

    test("should return 403 if logged user's role is not correct", async () => {
      token = (
        await supertest(app)
          .post('/v1/auth/login')
          .send({ username: newCandidate.username, password: newCandidate.password })
          .expect(200)
      ).body.data
      const res = await supertest(app).post(`/v1/post/:${postId}`).set('Authorization', `Bearer ${token}`).send(payload)

      const user = await findUserByField('username', newCandidate.username)
      if (user) {
        const candidate = await findCandidateByField('userId', user._id.toString())
        if (candidate) {
          const post = await findPostByField('candidateId', candidate._id.toString())
          if (post) {
            expect(res.body.status).toBe(403)
            expect(candidate.posts).not.toStrictEqual([post._id.toString()])
            expect(user.comments).not.toStrictEqual([post._id.toString()])
          }
        }
      }
    })
  })
})
