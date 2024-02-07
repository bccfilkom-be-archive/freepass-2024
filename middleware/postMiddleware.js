const pool = require('../config/database');

const checkPostExistence = (req, res, next) => {
  const { id, 'post-id': postId } = req.query;

  if (!id && !postId) {
    return next();
  } else {
    let targetId;

    if (id) {
      targetId = id;
    } else if (postId) {
      targetId = postId;
    }  
  
    pool.query(`SELECT * FROM post WHERE id = ?`, [targetId], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }
  
      if (results.length === 0) {
        return res.status(404).json({ error: 'Post not found' });
      }
  
      next();
    });
  }
};


const checkPostOwnership = (req, res, next) => {
  const postId = req.query.id;
  const userId = req.session.userId;

  if (req.session.status === 'admin') {
    return next();
  }

  pool.query(`SELECT user_id FROM post WHERE id = ?`, [postId], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    if (results.length === 0 || results[0].user_id !== userId) {
      return res.status(403).json({ error: 'Forbidden - You do not have permission to edit or delete this post' });
    }

    next();
  });
};

module.exports = {
  checkPostExistence,
  checkPostOwnership
};