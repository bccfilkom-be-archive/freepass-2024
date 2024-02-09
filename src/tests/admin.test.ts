import supertest from 'supertest'
import { MongoMemoryServer } from 'mongodb-memory-server'
import mongoose from 'mongoose'
import { app } from '../app'
import { User } from '../models/user.model'
import type { RegisterForm } from '../types/auth.type'
import { hashing } from '../utils/bcrypt'
import { findUserByField, findUserById } from '../services/user.service'
import { findCandidateByField } from '../services/candidate.service'
import { findPostByField, findPostById } from '../services/post.service'
import { findCommentById } from '../services/comment.service'

describe('adminRoutes', () => {
  beforeAll(async () => {
    const mongoServer = await MongoMemoryServer.create()
    await mongoose.connect(mongoServer.getUri())
  })

  afterAll(async () => {
    await mongoose.disconnect()
    await mongoose.connection.close()
  })

  describe('post /v1/admin/:id', () => {
    let newUser: RegisterForm
    beforeAll(async () => {
      newUser = {
        fullName: 'valid full name',
        username: 'validusername',
        nim: '231502001110111',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newUser)

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
    })

    test("should return 200 if id and logged user's role is correct", async () => {
      const token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin', password: 'password' })).body
        .data

      let user = await findUserByField('username', newUser.username)

      const res = await supertest(app).post(`/v1/admin/${user?._id}`).set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(200)

      user = await findUserByField('username', newUser.username)
      expect(user?.role).toEqual('candidate')

      if (user) {
        const candidate = await findCandidateByField('userId', user._id.toString())
        expect(candidate).toBeDefined()
      }
    })

    test('should return 400 if id is wrong', async () => {
      const token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin', password: 'password' })).body
        .data

      const res = await supertest(app).post('/v1/admin/wrongid').set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(400)
    })

    test("should return 403 if logged user's role is wrong", async () => {
      const token = (
        await supertest(app).post('/v1/auth/login').send({ username: newUser.username, password: newUser.password })
      ).body.data

      const res = await supertest(app).post('/v1/admin/wrongid').set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(403)
    })
  })

  describe('delete /v1/admin/user/:id', () => {
    let newUser: RegisterForm
    let newCandidate: RegisterForm
    let token: string
    beforeAll(async () => {
      newUser = {
        fullName: 'valid full name',
        username: 'validusername2',
        nim: '231502001110112',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail2@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newUser)

      newCandidate = {
        fullName: 'valid full name',
        username: 'validusername3',
        nim: '231502001110113',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail3@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newCandidate)

      const password = hashing('password')
      const admin = new User({
        fullName: 'admin',
        nim: '0000002',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'admin2@gmail.com',
        username: 'admin2',
        password,
        role: 'admin'
      })
      await admin.save()

      const userCandidate = await findUserByField('username', newCandidate.username)
      if (userCandidate) {
        token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin2', password: 'password' })).body
          .data
        await supertest(app).post(`/v1/admin/${userCandidate._id}`).set('Authorization', `Bearer ${token}`)

        const candidate = await findCandidateByField('userId', userCandidate._id.toString())
        if (candidate) {
          token = (
            await supertest(app)
              .post('/v1/auth/login')
              .send({ username: newCandidate.username, password: newCandidate.password })
          ).body.data
          let payload = {
            caption: 'valid caption'
          }
          await supertest(app).post('/v1/post').set('Authorization', `Bearer ${token}`).send(payload).expect(201)

          token = (
            await supertest(app).post('/v1/auth/login').send({ username: newUser.username, password: newUser.password })
          ).body.data
          payload = {
            caption: 'valid caption'
          }
          const post = await findPostByField('candidateId', candidate._id.toString())
          if (post) {
            await supertest(app)
              .post(`/v1/post/${post._id}`)
              .set('Authorization', `Bearer ${token}`)
              .send(payload)
              .expect(201)
          }
        }
      }
    })

    test("should return 200 if id is ok and logged user's role is correct (delete user)", async () => {
      token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin2', password: 'password' })).body.data
      const user = await findUserByField('username', newUser.username)
      if (user) {
        const commentId = user.comments.at(0) ?? ''
        const comment = await findCommentById(commentId.toString())
        if (comment) {
          const postId = comment.postId
          const post = await findPostById(postId.toString())
          if (post) {
            const res = await supertest(app)
              .delete(`/v1/admin/user/${user._id}`)
              .set('Authorization', `Bearer ${token}`)
            expect(res.body.status).toBe(200)
            const checkPost = await findPostById(postId.toString())
            expect(checkPost?.comments).toHaveLength(0)
            const checkComment = await findCommentById(commentId.toString())
            expect(checkComment).toBe(null)
            const checkUser = await findUserByField('username', newUser.username)
            expect(checkUser).toBe(null)
            const checkCommentedUser = await findUserById(comment.userId.toString())
            if (checkCommentedUser) {
              expect(checkCommentedUser.comments).toHaveLength(0)
              expect(checkCommentedUser.commentedPosts).toHaveLength(0)
            }
          }
        }
      }
    })

    test("should return 200 if id is ok and logged user's role is correct (delete user who is candidate)", async () => {
      const user = await findUserByField('username', newCandidate.username)
      if (user) {
        const candidate = await findCandidateByField('userId', user._id.toString())
        if (candidate) {
          await supertest(app).post('/v1/auth/register').send(newUser)
          token = (
            await supertest(app).post('/v1/auth/login').send({ username: newUser.username, password: newUser.password })
          ).body.data
          const payload = {
            caption: 'valid caption'
          }
          const newPost = await findPostByField('candidateId', candidate._id.toString())
          if (newPost) {
            await supertest(app)
              .post(`/v1/post/${newPost._id}`)
              .set('Authorization', `Bearer ${token}`)
              .send(payload)
              .expect(201)
          }

          token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin2', password: 'password' })).body
            .data

          const postId = candidate.posts.at(0) ?? ''
          const post = await findPostById(postId.toString())
          if (post) {
            const comments = post.comments.at(0)

            let commentId = ''
            if (comments) {
              commentId = comments._id.toString()
            }
            const comment = await findCommentById(commentId)
            if (comment) {
              const res = await supertest(app)
                .delete(`/v1/admin/user/${user._id}`)
                .set('Authorization', `Bearer ${token}`)
              expect(res.body.status).toBe(200)
              const checkComment = await findCommentById(commentId)
              expect(checkComment).toBe(null)
              const checkPost = await findPostById(postId.toString())
              expect(checkPost).toBe(null)
              const checkCandidate = await findCandidateByField('userId', user._id.toString())
              expect(checkCandidate).toBe(null)
              const checkUser = await findUserByField('username', newCandidate.username)
              expect(checkUser).toBe(null)
            }
          }
        }
      }
    })
  })

  describe('delete /v1/admin/candidate/:id', () => {
    let newUser: RegisterForm
    let newCandidate: RegisterForm
    let token: string
    beforeAll(async () => {
      newUser = {
        fullName: 'valid full name',
        username: 'validusername4',
        nim: '231502001110114',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail4@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newUser)

      newCandidate = {
        fullName: 'valid full name',
        username: 'validusername5',
        nim: '231502001110115',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'validemail5@gmail.com',
        password: 'validpassword'
      }
      await supertest(app).post('/v1/auth/register').send(newCandidate)

      const password = hashing('password')
      const admin = new User({
        fullName: 'admin',
        nim: '0000003',
        fakultas: 'valid fakultas',
        prodi: 'valid prodi',
        email: 'admin3@gmail.com',
        username: 'admin3',
        password,
        role: 'admin'
      })
      await admin.save()

      const userCandidate = await findUserByField('username', newCandidate.username)
      if (userCandidate) {
        token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin3', password: 'password' })).body
          .data
        await supertest(app).post(`/v1/admin/${userCandidate._id}`).set('Authorization', `Bearer ${token}`)

        const candidate = await findCandidateByField('userId', userCandidate._id.toString())
        if (candidate) {
          token = (
            await supertest(app)
              .post('/v1/auth/login')
              .send({ username: newCandidate.username, password: newCandidate.password })
          ).body.data
          let payload = {
            caption: 'valid caption'
          }
          await supertest(app).post('/v1/post').set('Authorization', `Bearer ${token}`).send(payload).expect(201)

          token = (
            await supertest(app).post('/v1/auth/login').send({ username: newUser.username, password: newUser.password })
          ).body.data
          payload = {
            caption: 'valid caption'
          }
          const post = await findPostByField('candidateId', candidate._id.toString())
          if (post) {
            await supertest(app)
              .post(`/v1/post/${post._id}`)
              .set('Authorization', `Bearer ${token}`)
              .send(payload)
              .expect(201)
          }
        }
      }
    })

    test("should return 200 if id is ok and logged user's role is correct (delete user who is candidate)", async () => {
      const user = await findUserByField('username', newCandidate.username)
      if (user) {
        const candidate = await findCandidateByField('userId', user._id.toString())
        if (candidate) {
          token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin3', password: 'password' })).body
            .data

          const postId = candidate.posts.at(0) ?? ''
          const post = await findPostById(postId.toString())
          if (post) {
            const comments = post.comments.at(0)

            let commentId = ''
            if (comments) {
              commentId = comments._id.toString()
            }
            const comment = await findCommentById(commentId)
            if (comment) {
              const res = await supertest(app)
                .delete(`/v1/admin/candidate/${candidate._id}`)
                .set('Authorization', `Bearer ${token}`)
              expect(res.body.status).toBe(200)
              const checkComment = await findCommentById(commentId)
              expect(checkComment).toBe(null)
              const checkPost = await findPostById(postId.toString())
              expect(checkPost).toBe(null)
              const checkCandidate = await findCandidateByField('userId', user._id.toString())
              expect(checkCandidate).toBe(null)
              const checkUser = await findUserByField('username', newCandidate.username)
              expect(checkUser).toBe(null)
              const checkCommentedUser = await findUserById(comment.userId.toString())
              if (checkCommentedUser) {
                expect(checkCommentedUser.comments).toHaveLength(0)
                expect(checkCommentedUser.commentedPosts).toHaveLength(0)
              }
            }
          }
        }
      }
    })
  })
})
