import Joi from 'joi'
import type { CreateCommentForm, UpdateCommentForm } from '../types/comment.type'

export const createFormValidation = (payload: CreateCommentForm) => {
  const schema = Joi.object({
    caption: Joi.string().required()
  })

  return schema.validate(payload)
}

export const updateFormValidation = (payload: UpdateCommentForm) => {
  const schema = Joi.object({
    caption: Joi.string().required()
  })

  return schema.validate(payload)
}
