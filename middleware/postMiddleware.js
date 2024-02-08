const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const checkPostExistence = (req, res, next) => {
  const { id } = req.params;
  const { 'post_id': postId } = req.body;

  if (!id && !postId) {
    return next();
  } else {
    let targetId;

    if (id) {
      targetId = id;
    } else if (postId) {
      targetId = postId;
    }

    executeQuery(`SELECT * FROM post WHERE id = ?`, [targetId])
      .then((results) => {
        if (results.length === 0) {
          return res.status(404).json({ error: 'Post not found!' });
        } else {
          next();
        }
      })
      .catch((error) => {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      })
  }
};


const checkPostOwnership = (req, res, next) => {
  const postId = req.params.id;
  const userId = req.session.userId;

  if (req.session.status === 'admin') {
    return next();
  }

  executeQuery(`SELECT user_id FROM post WHERE id = ?`, [postId])
    .then((results) => {
      if (results.length === 0 || results[0].user_id !== userId) {
        return res.status(403).json({ error: 'You do not have permission to edit or delete this post!' });
      } else {
        next();
      }
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
};

module.exports = {
  checkPostExistence,
  checkPostOwnership
};