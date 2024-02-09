import { Vote } from '../models/vote.model'
import type { CreateVoteForm } from '../types/vote.type'

export const createVoteService = async (payload: CreateVoteForm, hashedUserId: string) => {
  return await Vote.create({ ...payload, hashedUserId })
}

export const getAllVotes = async () => {
  return await Vote.find({})
}

export const findVoteById = async (id: string) => {
  return await Vote.findById(id)
}
