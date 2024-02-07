import supertest from 'supertest'
import { MongoMemoryServer } from 'mongodb-memory-server'
import mongoose from 'mongoose'
import { app } from '../app'
import { User } from '../models/user.model'
import type { RegisterForm } from '../types/auth.type'
import { hashing } from '../utils/bcrypt'
import { findUserByField } from '../services/user.service'

describe('candidateRoutes', () => {
  beforeAll(async () => {
    const mongoServer = await MongoMemoryServer.create()
    await mongoose.connect(mongoServer.getUri())
  })

  afterAll(async () => {
    await mongoose.disconnect()
    await mongoose.connection.close()
  })

  describe('get /v1/candidate', () => {
    let newUser: RegisterForm
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

      token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin2', password: 'password' })).body.data
      const user = await findUserByField('username', newUser.username)
      if (user) await supertest(app).post(`/v1/admin/${user._id}`).set('Authorization', `Bearer ${token}`)
    })

    test("should return 200 if logged user's role correct", async () => {
      const res = await supertest(app).get('/v1/candidate').set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(200)
    })

    test("should return 403 if logged user's role not correct", async () => {
      const token = (
        await supertest(app).post('/v1/auth/login').send({ username: newUser.username, password: newUser.password })
      ).body.data
      const res = await supertest(app).get('/v1/candidate').set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(403)
    })
  })

  describe('get /v1/user/:id', () => {
    let newUser: RegisterForm
    let token: string
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

      token = (await supertest(app).post('/v1/auth/login').send({ username: 'admin', password: 'password' })).body.data
    })

    test("should return 200 if id is ok and logged user's role correct", async () => {
      const user = await findUserByField('username', newUser.username)
      if (user) {
        await supertest(app).post(`/v1/admin/${user._id}`).set('Authorization', `Bearer ${token}`)
        const res = await supertest(app).get(`/v1/user/${user._id}`).set('Authorization', `Bearer ${token}`)
        expect(res.body.status).toBe(200)
      }
    })

    test('should return 400 if id is not ok', async () => {
      const userId = 'wronguserid'
      const res = await supertest(app).get(`/v1/user/${userId}`).set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(400)
    })

    test("should return 403 if logged user's role not correct", async () => {
      const token = (
        await supertest(app).post('/v1/auth/login').send({ username: newUser.username, password: newUser.password })
      ).body.data
      const user = await findUserByField('username', newUser.username)
      if (user) {
        await supertest(app).post(`/v1/admin/${user._id}`).set('Authorization', `Bearer ${token}`)
        const res = await supertest(app).get(`/v1/user/${user._id}`).set('Authorization', `Bearer ${token}`)
        expect(res.body.status).toBe(403)
      }
    })
  })
})
