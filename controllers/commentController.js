const pool = require('../config/database');

const addComment = (req, res) => {
  const { postId, content } = req.body;
  pool.query(`INSERT INTO comment (user_id, post_id, content) VALUES (?, ?, ?)`, [req.session.userId, postId, content], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Comment sent successfully' });
  });
};

const deleteComment = (req, res) => {
  const { id } = req.query;
  pool.query(`DELETE FROM comment WHERE id = ?`, [id], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Comment deleted successfully' });
  });
};

module.exports = {
  addComment,
  deleteComment
};
