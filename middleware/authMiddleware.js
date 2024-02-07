const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const authenticateUser = (req, res, next) => {
  if (req.session.loggedin) {
    next();
  } else {
    return res.status(401).json({ error: 'Please log in' });
  }
};

const checkUsername = (req, res, next) => {
  const { username } = req.body;

  if (!username || (req.session.username == username)) {
    return next();
  }

  executeQuery(`SELECT * FROM user WHERE username = ?`, [username])
    .then((results) => {
      if (results.length === 0) {
        next();
      } else {
        return res.status(409).json({ error: 'Username already used!' });
      }
    })
    .catch((error) => {
      console.log(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
}

const checkLogOut = (req, res, next) => {
  if (req.session.loggedin) {
    return res.status(401).json({ error: 'You must logout before accessing this route!' });
  } else {
    next();
  }
}

module.exports = {
  authenticateUser,
  checkUsername,
  checkLogOut
};
