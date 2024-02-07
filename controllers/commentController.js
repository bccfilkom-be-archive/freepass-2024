const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const viewComment = (req, res) => {
  const { id, 'post-id': postId, username } = req.query;

  let query = 'SELECT * FROM comment WHERE';
  let values = [];

  if (id) {
    query += ' id = ?';
    values.push(id);
  } else if (postId) {
    query += ' post_id = ?';
    values.push(postId);
  } else if (username) {
    query += ' user_id IN (SELECT id FROM user WHERE username = ?)';
    values.push(username);
  } else {
    return res.status(400).json({ error: 'Provide id, postId, or username!' });
  }

  executeQuery(query, values)
    .then((results) => {
      if (results.length === 0) {
        res.status(400).json({ error: 'No results' });
      } else {
        res.json({ results });
      }
    })
    .catch((error) => {
      console.log(error);
      res.status(500).json({ error: 'Internal Server Error' });
    });
};

const addComment = (req, res) => {
  const { 'post-id': postId } = req.query;
  const { content } = req.body;

  const query = `INSERT INTO comment (user_id, post_id, content) VALUES (?, ?, ?)`;
  const values = [req.session.userId, postId, content];

  executeQuery(query, values)
    .then(() => {
      res.json({ message: 'Comment sent successfully' });
    })
    .catch((error) => {
      console.error(error);
      res.status(500).json({ error: 'Internal Server Error' });
    });
};

const deleteComment = (req, res) => {
  const { id } = req.query;

  const query = `DELETE FROM comment WHERE id = ?`;
  const values = [id];

  executeQuery(query, values)
    .then(() => {
      res.json({ message: 'Comment deleted successfully' });
    })
    .catch((error) => {
      console.error(error);
      res.status(500).json({ error: 'Internal Server Error' });
    });
};

module.exports = {
  viewComment,
  addComment,
  deleteComment
};
