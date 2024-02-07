import Joi from 'joi'
import type { CreatePostForm } from '../types/post.type'

export const createFormValidation = (payload: CreatePostForm) => {
  const schema = Joi.object({
    caption: Joi.string().required()
  })

  return schema.validate(payload)
}
