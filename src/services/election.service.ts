import { Election } from '../models/election.model'
import type { CreateElectionForm } from '../types/election.type'

export const createElectionService = async (payload: CreateElectionForm) => {
  return await Election.create(payload)
}
