const pool = require('../config/database');

const viewPost = (req, res) => {
  const { id, username } = req.query;

  if (id && username) {
    return res.status(400).json({ error: 'Provide either id or username, not both' });
  }

  function fetchCommentsForPost(post) {
    return new Promise((resolve, reject) => {
      pool.query(`SELECT c.id, c.user_id, u.username AS username, u.name AS name, c.post_id, c.content, c.timestamp
        FROM comment AS c
        JOIN user u ON user_id = u.id
        WHERE c.post_id = ?`,
        [post.id],
        (error, commentResults) => {
          if (error) {
            reject(error);
          } else {
            post.comments = commentResults;
            resolve();
          }
        });
    });
  }

  let baseQuery = `SELECT p.id, p.user_id, u.username AS username, u.name AS name, p.title, p.content, p.timestamp 
    FROM post AS p
    JOIN user u ON user_id = u.id`;

  if (id) {
    baseQuery += ` WHERE p.id = ${id}`;
  } else if (username) {
    baseQuery += ` WHERE p.user_id IN (SELECT id FROM user WHERE username = '${username}')`;
  }

  pool.query(baseQuery, (error, posts) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    if (posts.length === 0) {
      return res.status(400).json({ message: 'No result' });
    }

    const promises = posts.map((post) => {
      return fetchCommentsForPost(post);
    });

    Promise.all(promises)
      .then(() => {
        res.json(posts);
      })
      .catch((error) => {
        console.error(error);
        res.status(500).json({ error: 'Internal Server Error' });
      });
  });
};

const addPost = (req, res) => {
  const { title, content } = req.body;
  pool.query(`INSERT INTO post (user_id, title, content) VALUES (?, ?, ?)`, [req.session.userId, title, content], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Post sent successfully' });
  });
};

const editPost = (req, res) => {
  const { id } = req.query;
  const { title, content } = req.body;

  pool.query(`UPDATE post SET title = ?, content = ? WHERE id = ?`, [title, content, id], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Post updated successfully' });
  });
};

const deletePost = (req, res) => {
  const { id } = req.query;

  if (!id) {
    return res.status(400).json({ error: 'Provide post id' });
  }

  pool.query(`DELETE FROM post WHERE id = ?`, [id], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    res.json({ message: 'Post deleted successfully' });
  });
};

module.exports = {
  viewPost,
  addPost,
  editPost,
  deletePost
};
