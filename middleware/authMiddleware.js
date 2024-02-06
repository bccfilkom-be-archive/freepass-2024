const pool = require('../config/database');

const authenticateUser = (req, res, next) => {
  if (req.session.loggedin) {
    next();
  } else {
    res.status(401).json({ error: 'Unauthorized - Please log in' });
  }
};

module.exports = {
  authenticateUser,
};
