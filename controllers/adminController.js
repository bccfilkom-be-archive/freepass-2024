const pool = require('../config/database');

const editStatus = (req, res) => {
  const { username, status } = req.body;
  
  if (username == req.session.username) {
    req.session.status = status;
  }

  pool.query(`UPDATE user SET status = ? WHERE username = ?`, [status, username], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'User status updated successfully' });
  });
};

module.exports = {
  editStatus
};