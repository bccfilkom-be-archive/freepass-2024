const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const checkCommentExistence = (req, res, next) => {
  const { id } = req.query;

  executeQuery(`SELECT * FROM comment WHERE id = ?`, [id])
  .then((results) => {
    if (results.length === 0) {
      return res.status(404).json({ error: 'Comment not found' });
    } else {
      next();
    }
  })
  .catch((error) => {
    console.error(error);
    return res.status(500).json({ error: 'Internal Server Error' });
  });
};

module.exports = {
  checkCommentExistence
};