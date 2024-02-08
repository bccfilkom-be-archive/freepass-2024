const bcrypt = require('bcrypt');
const { executeQuery } = require('../services/db');
const pool = require('../config/database');

const register = (req, res) => {
  const { nim, name, username, password, major, faculty, status, description } = req.body;

  if (!username || !password) {
    return res.status(400).json({ error: 'Provide at least username and password!' });
  }

  if (status != 'user' && status != 'admin' && status != 'candidate') {
    return res.status(400).json({ error: 'Provide a valid status!' });
  }

  bcrypt.hash(password, 10, (err, hash) => {
    if (err) {
      console.error(err);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    let query = `INSERT INTO user (nim, name, username, password, major, faculty, status, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`;
    let values = [nim, name, username, hash, major, faculty, status ? status : 'user', description];

    executeQuery(query, values)
      .then(() => {
        res.json({ message: 'User registered successfully', username });
      })
      .catch((error) => {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      });
});
};

const login = (req, res) => {
  const { username, password } = req.body;

  if (username && password) {
    const query = `SELECT * FROM user WHERE username = ?`;
    const values = [username];

    executeQuery(query, values)
      .then((results) => {
        if (results.length > 0) {
          const storedHashedPassword = results[0].password;
          const status = results[0].status;
          const userId = results[0].id;

          bcrypt.compare(password, storedHashedPassword, (err, passwordMatch) => {
            if (err || !passwordMatch) {
              return res.status(401).json({ error: 'Incorrect Username and/or Password!' });
            }

            req.session.loggedin = true;
            req.session.username = username;
            req.session.status = status;
            req.session.userId = userId;

            res.json({ message: 'User logged in successfully', username, status });
          });
        } else {
          res.status(401).json({ error: 'Incorrect Username and/or Password!' });
        }
      })
      .catch((error) => {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      });
  } else {
    res.status(400).json({ error: 'Please enter Username and Password!' });
  }
};

const logout = (req, res) => {
  req.session.destroy((err) => {
    if (err) {
      console.error(err);
      res.status(500).json({ error: 'Internal Server Error' });
    } else {
      res.json({ message: 'Logout successful' });
    }
  });
};

module.exports = {
  register,
  login,
  logout,
};
