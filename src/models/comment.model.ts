import mongoose from 'mongoose'

const commentSchema = new mongoose.Schema(
  {
    caption: {
      type: String,
      required: true,
      trim: true
    },
    postId: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'Post',
      required: true
    },
    userId: {
      type: mongoose.Schema.Types.ObjectId,
      ref: 'User',
      required: true
    }
  },
  {
    timestamps: true
  }
)

commentSchema.set('toJSON', {
  transform: (document, returnedObject) => {
    returnedObject.id = returnedObject._id.toString()
    delete returnedObject._id
    delete returnedObject.__v
  }
})

/**
 * @typedef Post
 */
export const Comment = mongoose.model('Comment', commentSchema)
