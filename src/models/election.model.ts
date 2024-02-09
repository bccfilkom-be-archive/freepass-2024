import mongoose from 'mongoose'

const electionSchema = new mongoose.Schema(
  {
    startDate: {
      type: Date,
      required: true
    },
    endDate: {
      type: Date,
      required: true
    },
    votes: [
      {
        type: mongoose.Schema.Types.ObjectId,
        ref: 'Vote'
      }
    ]
  },
  {
    timestamps: true
  }
)

electionSchema.set('toJSON', {
  transform: (document, returnedObject) => {
    returnedObject.id = returnedObject._id.toString()
    delete returnedObject._id
    delete returnedObject.__v
  }
})

/**
 * @typedef Election
 */
export const Election = mongoose.model('Election', electionSchema)
