const pool = require('../config/database');

const checkUserExistence = (req, res, next) => {
  const { username, id } = req.query;

  if (!username && !id) {
    return next();
  }

  if (username) {
    pool.query(`SELECT * FROM user WHERE username = ?`, [username], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (results.length === 0) {
        return res.status(404).json({ error: 'User not found!' });
      }

      next();
    });
  } else if (id) {
    pool.query(`SELECT * FROM user WHERE id = ?`, [id], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (results.length === 0) {
        return res.status(404).json({ error: 'User not found!' });
      }

      next();
    });
  }
};

const checkUserStatus = (requiredStatus) => {
  return (req, res, next) => {
    if (!requiredStatus.includes(req.session.status)) {
      return res.status(403).json({
        error: `Insufficient privileges, only ${requiredStatus.join(' or ')} can access`,
      });
    }
    next();
  };
};

module.exports = {
  checkUserExistence,
  checkUserStatus
};