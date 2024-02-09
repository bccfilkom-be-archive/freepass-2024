import { Election } from '../models/election.model'
import type { CreateElectionForm } from '../types/election.type'

export const createElectionService = async (payload: CreateElectionForm) => {
  return await Election.create(payload)
}

export const getAllElections = async () => {
  return await Election.find({})
}

export const findElectionById = async (id: string) => {
  return await Election.findById(id)
}
