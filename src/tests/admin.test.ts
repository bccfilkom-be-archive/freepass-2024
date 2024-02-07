import supertest from 'supertest'
import { MongoMemoryServer } from 'mongodb-memory-server'
import mongoose from 'mongoose'
import { app } from '../app'
import { User } from '../models/user.model'
import type { RegisterForm } from '../types/auth.type'
import { hashing } from '../utils/bcrypt'

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

      let user = await User.findOne({ username: newUser.username })

      const res = await supertest(app).post(`/v1/admin/${user?._id}`).set('Authorization', `Bearer ${token}`)
      expect(res.body.status).toBe(200)

      user = await User.findOne({ username: newUser.username })
      expect(user?.role).toEqual('candidate')
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
})
