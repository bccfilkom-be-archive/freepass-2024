const bcrypt = require('bcrypt');
const pool = require('../config/database');

const register = (req, res) => {
  const { nim, name, username, password, major, faculty, status, description } = req.body;

  if (!username || !password) {
    return res.status(400).json({ error: 'Provide at least username and password' });
  }

  bcrypt.hash(password, 10, (err, hash) => {
    if (err) {
      console.error(err);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    pool.query(`INSERT INTO user (nim, name, username, password, major, faculty, status, description) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
      [nim, name, username, hash, major, faculty, status ? status : 'user', description], (error, results) => {
        if (error) {
          console.error(error);
          return res.status(500).json({ error: 'Internal Server Error' });
        }
        res.json({ message: 'User registered successfully', username });
      });
  });
};

const login = (req, res) => {
  const { username, password } = req.body;

  if (username && password) {
    pool.query(`SELECT * FROM user WHERE username = ?`, [username], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

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

          res.json({ message: 'User logged in successfully', username });
        });
      } else {
        res.status(401).json({ error: 'Incorrect Username and/or Password!' });
      }
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
