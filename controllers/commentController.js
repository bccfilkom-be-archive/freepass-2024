const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const viewComment = (req, res) => {
  const { id, 'post-id': postId, username } = req.query;

  let query = `
  SELECT c.id, c.user_id, u.username, u.name, c.post_id, c.content, c.timestamp
  FROM comment AS c
  JOIN user AS u ON c.user_id = u.id
  WHERE 1
  `;
  let values = [];

  if (id) {
    query += ' AND c.id = ?';
    values.push(id);
  }

  if (postId) {
    query += ' AND c.post_id = ?';
    values.push(postId);
  }

  if (username) {
    query += ' AND c.user_id = (SELECT id FROM user WHERE username = ?)';
    values.push(username);
  }

  executeQuery(query, values)
    .then((results) => {
      if (results.length === 0) {
        return res.status(400).json({ error: 'No results' });
      } else {
        res.json(results);
      }
    })
    .catch((error) => {
      console.log(error);
      res.status(500).json({ error: 'Internal Server Error' });
    });
};

const addComment = (req, res) => {
  const { content, post_id: postId } = req.body;

  const query = `INSERT INTO comment (user_id, post_id, content) VALUES (?, ?, ?)`;
  const values = [req.session.userId, postId, content];

  if (!postId || !content) {
    return res.status(400).json({ error: 'Provide post_id and content in the request body!' });
  }

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
  const { id } = req.params;

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
