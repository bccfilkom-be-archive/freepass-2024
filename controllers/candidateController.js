const pool = require('../config/database');

const viewCandidate = (req, res) => {
  const { username } = req.query;

  pool.query(`SELECT * FROM user WHERE username = ? AND status = 'candidate'`, [username == null ? req.session.username : username], (userError, userResults) => {
    if (userError) {
      console.error(userError);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    if (userResults.length === 0) {
      return res.status(400).json({ message: 'Candidate not found!' });
    }

    const user = userResults[0];

    pool.query(`SELECT * FROM post WHERE user_id IN (SELECT id FROM user WHERE username = ?)`, [username], (error, posts) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (posts.length === 0) {
        return res.json({profile: user, posts: {}});
      }

      posts.map((post, index) => {
        pool.query(`SELECT * FROM comment WHERE post_id = ?`, [post.id], (commentError, commentResults) => {
          if (commentError) {
            console.error(commentError);
            return res.status(500).json({ error: 'Internal Server Error' });
          }

          post.comments = commentResults;

          if (index === posts.length - 1) {
            const candidateData = {
              profile: user,
              posts: posts
            };
            res.json(candidateData);
          }
        });
      });
    });
  });
};

module.exports = {
  viewCandidate
}
