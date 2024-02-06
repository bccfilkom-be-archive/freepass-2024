import Joi from 'joi'
import type { UserType } from '../types/user.type'

export const createUserValidation = (payload: UserType) => {
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
