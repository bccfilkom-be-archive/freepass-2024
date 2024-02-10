const pool = require('../config/database');
const { executeQuery } = require('../services/db');
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
    next();
  } catch (error) {
    return res.status(401).json({ error: 'Unauthorized: Invalid token' });
  }

};

const getUserInfo = async (req, res, next) => {
  try {
    const id = req.user.userId;
    const results = await executeQuery(`SELECT * FROM user WHERE id = ?`, [id]);
    req.user.status = results[0].status;
    req.user.username = results[0].username;
    next();
  } catch (error) {
    console.error(error);
    return res.status(500).json({ error: 'Internal Server Error' });
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
  getUserInfo
};
