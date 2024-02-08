const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const checkUserExistence = (req, res, next) => {
  const { username, id } = req.params;

  if (!username && !id) {
    return next();
  }

  let query = 'SELECT * FROM user WHERE';
  let values = [];

  if (username) {
    query += ' username = ?';
    values.push(username);
  } else if (id) {
    query += ' id = ?';
    values.push(id);
  }

  executeQuery(query, values)
    .then((results) => {
      if (results.length === 0) {
        return res.status(404).json({ error: 'User not found!' });
      } else {
        next();
      }
    })
    .catch((error) => {
      console.error(error);
      res.status(500).json({ error: 'Internal Server Error' });
    });
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