import Joi from 'joi'
import type { UpdateUserForm } from '../types/user.type'

export const updateUserValidation = (payload: UpdateUserForm) => {
  const schema = Joi.object({
    fullName: Joi.string(),
    nim: Joi.string(),
    fakultas: Joi.string(),
    prodi: Joi.string(),
    username: Joi.string(),
    email: Joi.string().email(),
    password: Joi.string()
  })

  return schema.validate(payload)
}
