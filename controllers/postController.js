const pool = require('../config/database');

const viewPost = (req, res) => {
  const { id, username } = req.query;

  if (id && username) {
    return res.status(400).json({ error: 'Provide either id or username, not both' });
  }

  if (!id && !username) {
    pool.query(`SELECT * FROM post`, (error, posts) => {
      if (error) throw error;

      if (posts.length === 0) {
        return res.status(400).json({ message: 'No result' });
      }

      posts.map((post, index) => {
        pool.query(`SELECT * FROM comment WHERE post_id = ?`, [post.id], (commentError, commentResults) => {
          if (commentError) throw commentError;

          post.comments = commentResults;

          if (index === posts.length - 1) {
            res.json(posts);
          }
        });
      });
    });
  }

  if (id) {
    pool.query(`SELECT * FROM post WHERE id = ?`, [id], (error, posts) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (posts.length === 0) {
        return res.status(400).json({ message: 'No result' });
      }

      const post = posts[0];
      pool.query(`SELECT * FROM comment WHERE post_id = ?`, [id], (commentError, commentResults) => {
        if (commentError) {
          console.error(commentError);
          return res.status(500).json({ error: 'Internal Server Error' });
        }

        post.comments = commentResults;
        res.json(post);
      });
    });
  }

  if (username) {
    pool.query(`SELECT * FROM post WHERE user_id IN (SELECT id FROM user WHERE username = ?)`, [username], (error, posts) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (posts.length === 0) {
        return res.status(400).json({ message: 'No result' });
      }

      posts.map((post, index) => {
        pool.query(`SELECT * FROM comment WHERE post_id = ?`, [post.id], (commentError, commentResults) => {
          if (commentError) {
            console.error(commentError);
            return res.status(500).json({ error: 'Internal Server Error' });
          }

          post.comments = commentResults;

          if (index === posts.length - 1) {
            res.json(posts);
          }
        });
      });
    });
  }
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
