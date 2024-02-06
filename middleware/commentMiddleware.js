const pool = require('../config/database');

const checkCommentExistence = (req, res, next) => {
  const { id } = req.query;

  pool.query(`SELECT * FROM comment WHERE id = ?`, [id], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    if (results.length === 0) {
      return res.status(404).json({ error: 'Comment not found' });
    }

    next();
  });
};

module.exports = {
  checkCommentExistence
};