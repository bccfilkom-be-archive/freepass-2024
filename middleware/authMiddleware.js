const pool = require('../config/database');
const { executeQuery, getUserInfo } = require('../services/db');
const jwt = require('jsonwebtoken');

const authenticateUser = (req, res, next) => {
  const authHeader = req.headers.authorization;

  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    return res.status(401).json({ error: 'Unauthorized: Missing or invalid token' });
  }

  const token = authHeader.split(' ')[1];

  if (!token) {
    return res.status(401).json({ error: 'Unauthorized: Missing token' });
  }

  try {
    const decoded = jwt.verify(token, 'secret');
    req.user = decoded;
    getUserInfo(req.user.userId)
      .then(userInfo => {
        req.user.status = userInfo.status;
        req.user.username = userInfo.username;
        next();
      })
      .catch(error => {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      });
  } catch (error) {
    return res.status(401).json({ error: 'Unauthorized: Invalid token' });
  }

};

const checkUsername = (req, res, next) => {
  const { username } = req.body;

  if (!username) {
    return next();
  } else if (req.user) {
    if (req.user.username == username) {
      return next();
    }
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

module.exports = {
  authenticateUser,
  checkUsername,
};
