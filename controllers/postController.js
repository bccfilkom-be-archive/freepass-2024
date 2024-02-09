const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const viewPost = (req, res) => {
  const { id, username } = req.query;

  function fetchCommentsForPost(post) {
    return executeQuery(
      `SELECT c.id, c.user_id, u.username AS username, u.name AS name, c.post_id, c.content, c.timestamp
      FROM comment AS c
      JOIN user u ON user_id = u.id
      WHERE c.post_id = ?`,
      [post.id]
    );
  }

  let query = `SELECT p.id, p.user_id, u.username AS username, u.name AS name, p.title, p.content, p.timestamp 
    FROM post AS p
    JOIN user AS u ON p.user_id = u.id
    WHERE 1`;

  const values = [];

  if (id) {
    query += ` AND p.id = ?`;
    values.push(id);
  }
  if (username) {
    query += ` AND p.user_id = (SELECT id FROM user WHERE username = ?)`;
    values.push(username);
  }

  executeQuery(query, values)
    .then((posts) => {
      if (posts.length === 0) {
        return res.status(404).json({ error: 'Post(s) not found!' });
      }

      const promises = posts.map((post) => fetchCommentsForPost(post));

      Promise.all(promises)
        .then((commentsResults) => {
          commentsResults.forEach((commentResults, index) => {
            posts[index].comments = commentResults;
          });
          res.json(posts);
        })
        .catch((error) => {
          console.error(error);
          res.status(500).json({ error: 'Internal Server Error' });
        });
    })
    .catch((error) => {
      console.error(error);
      res.status(500).json({ error: 'Internal Server Error' });
    });
};

const addPost = (req, res) => {
  const { title, content } = req.body;

  let query = `INSERT INTO post (user_id, title, content) VALUES (?, ?, ?)`;
  let values = [req.user.userId, title, content];

  executeQuery(query, values)
    .then(() => {
      res.json({ message: 'Post sent successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
};

const editPost = (req, res) => {
  const { id } = req.params;
  const { title, content } = req.body;

  if (!id) {
    return res.status(400).json({ error: 'Provide post id!' });
  }

  let query = `UPDATE post SET title = ?, content = ? WHERE id = ?`;
  let values = [title, content, id];

  executeQuery(query, values)
    .then(() => {
      res.json({ message: 'Post updated successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
};

const deletePost = (req, res) => {
  const { id } = req.params;

  if (!id) {
    return res.status(400).json({ error: 'Provide post id!' });
  }

  executeQuery(`DELETE FROM post WHERE id = ?`, [id])
    .then(() => {
      res.json({ message: 'Post deleted successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
};

module.exports = {
  viewPost,
  addPost,
  editPost,
  deletePost
};
