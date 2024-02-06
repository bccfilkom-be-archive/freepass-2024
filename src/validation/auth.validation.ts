import Joi from 'joi'
import type { RegisterForm } from '../types/auth.type'

export const createUserValidation = (payload: RegisterForm) => {
  const schema = Joi.object({
    fullName: Joi.string().required(),
    nim: Joi.string().required(),
    fakultas: Joi.string().required(),
    prodi: Joi.string().required(),
    username: Joi.string().required(),
    email: Joi.string().required().email(),
    password: Joi.string().required()
  })

  return schema.validate(payload)
}
