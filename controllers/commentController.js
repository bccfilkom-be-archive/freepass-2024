const pool = require('../config/database');

const viewComment = (req, res) => {
  const { id, 'post-id': postId, username } = req.query;
  if (id) {
    pool.query(`SELECT * FROM comment WHERE id = ?`, [id], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }
      if (results.length === 0) {
        return res.status(400).json({ error: 'No results' });
      }
      res.json({ results });
    });
  } else if (postId) {
    pool.query(`SELECT * FROM comment WHERE post_id = ?`, [postId], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }
      if (results.length === 0) {
        return res.status(400).json({ error: 'No results' });
      }
      res.json({ results });
    });
  } else if (username) {
    pool.query(`SELECT * FROM comment WHERE user_id IN (SELECT id FROM user WHERE username = ?)`, [username], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }
      if (results.length === 0) {
        return res.status(400).json({ error: 'No results' });
      }
      res.json({ results });
    });
  } else {
    return res.status(400).json({ error: 'Provide id, postId, or username!' });
  }
};

const addComment = (req, res) => {
  const { 'post-id': postId } = req.query;
  const { content } = req.body;
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
  viewComment,
  addComment,
  deleteComment
};
