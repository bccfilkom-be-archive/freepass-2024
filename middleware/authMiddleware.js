const pool = require('../config/database');

const authenticateUser = (req, res, next) => {
  if (req.session.loggedin) {
    next();
  } else {
    return res.status(401).json({ error: 'Unauthorized - Please log in' });
  }
};

const checkUsername = (req, res, next) => {
  const { username } = req.body;

  if (!username) {
    return next();
  }

  pool.query(`SELECT * FROM user WHERE username = ?`, [username], (error, results) => {
    if (error) {
      console.log(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    
    if (results.length === 0) {
      next();
    } else {
      return res.status(400).json({ error: 'Username already used!' });
    }
  });
}

module.exports = {
  authenticateUser,
  checkUsername,
};
