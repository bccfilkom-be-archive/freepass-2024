import { Candidate } from '../models/candidate.model'

export const createCandidate = async (userId: string) => {
  return await Candidate.create({ userId })
}

export const findCandidateByField = async (field: string, value: string) => {
  return await Candidate.findOne({ [field]: value })
}

export const getAllCandidates = async () => {
  return await Candidate.find({}).populate('posts')
}

export const findCandidateById = async (id: string) => {
  return await Candidate.findById(id).populate('posts')
}
