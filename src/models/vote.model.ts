import mongoose from 'mongoose'

const voteSchema = new mongoose.Schema(
  {
    hashedUserId: {
      type: String,
      required: true
    },
    candidateId: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'Candidate',
      required: true
    }
  },
  {
    timestamps: true
  }
)

voteSchema.set('toJSON', {
  transform: (document, returnedObject) => {
    returnedObject.id = returnedObject._id.toString()
    delete returnedObject._id
    delete returnedObject.hashedUser
    delete returnedObject.__v
  }
})

/**
 * @typedef Vote
 */
export const Vote = mongoose.model('Vote', voteSchema)
