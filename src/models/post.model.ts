import mongoose from 'mongoose'

const postSchema = new mongoose.Schema(
  {
    caption: {
      type: String,
      required: true,
      trim: true
    },
    candidateId: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'Candidate',
      required: true
    },
    userIds: [
      {
        type: mongoose.Schema.Types.ObjectId,
        ref: 'User'
      }
    ]
  },
  {
    timestamps: true
  }
)

postSchema.set('toJSON', {
  transform: (document, returnedObject) => {
    returnedObject.id = returnedObject._id.toString()
    delete returnedObject._id
    delete returnedObject.__v
  }
})

/**
 * @typedef Post
 */
export const Post = mongoose.model('Post', postSchema)
