const bcrypt = require('bcrypt');
const pool = require('../config/database');

const viewAllUsers = (req, res) => {
  pool.query(`SELECT id, nim, username, name, major, faculty, status, description FROM user`, (error, results) => {
    if (error) throw error;
    res.json(results);
  });
};

const viewUser = (req, res) => {
  const { username } = req.query;

  const queryUser = (username) => {
    return new Promise((resolve, reject) => {
      pool.query(`SELECT id, nim, username, name, major, faculty, status, description FROM user WHERE username = ?`, [username], (error, results) => {
        if (error) {
          console.error(error);
          reject('Internal Server Error');
        } else {
          resolve(results);
        }
      });
    });
  };

  queryUser(username == null ? req.session.username : username)
    .then((results) => {
      if (results.length === 0) {
        res.status(404).json({ error: 'User not found' });
      } else {
        res.json(results);
      }
    })
    .catch((error) => {
      res.status(500).json({ error: error });
    });
};

const editProfile = (req, res) => {
  const { nim, name, username, password, major, faculty, description } = req.body;
  const updateFields = {};
  const updateValues = [];

  if (nim) {
    updateFields.nim = nim;
    updateValues.push(nim);
  }

  if (name) {
    updateFields.name = name;
    updateValues.push(name);
  }

  if (username) {
    updateFields.username = username;
    updateValues.push(username);
  }

  if (password) {
    bcrypt.hash(password, 10, (err, hash) => {
      if (err) {
        console.error(err);
        return res.status(500).json({ error: 'Internal Server Error' });
      }
      updateFields.password = hash;
      updateValues.push(hash);
    });
  }

  if (major) {
    updateFields.major = major;
    updateValues.push(major);
  }

  if (faculty) {
    updateFields.faculty = faculty;
    updateValues.push(faculty);
  }

  if (description) {
    updateFields.description = description;
    updateValues.push(description);
  }

  updateValues.push(req.session.userId);

  const updateQuery = `UPDATE user SET ` + Object.keys(updateFields).map(key => `${key} = ?`).join(`, `) + ` WHERE id = ?`;

  pool.query(updateQuery, updateValues, (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'User profile updated successfully' });
  });
};

const deleteUser = (req, res) => {
  const { id, username } = req.query;

  if (!id && !username) {
    return res.status(400).json({ error: 'Provide either id or username' });
  }

  if (id && username) {
    return res.status(400).json({ error: 'Provide either id or username, not both' });
  }

  if (id == req.session.userId || username == req.session.username) {
    return res.status(403).json({ error: 'You cannot delete your own account' });
  }

  if (id) {
    pool.query(`DELETE FROM user WHERE id = ?`, [id], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (results.affectedRows === 0) {
        return res.status(404).json({ error: 'User not found' });
      }

      res.json({ message: 'User deleted succesfully' })
    });
  }

  if (username) {
    pool.query(`DELETE FROM user WHERE username = ?`, [username], (error, results) => {
      if (error) {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (results.affectedRows === 0) {
        return res.status(404).json({ error: 'User not found' });
      }

      res.json({ message: 'User deleted succesfully' })
    });
  }
};

const editStatus = (req, res) => {
  const { username } = req.query;
  const { status } = req.body;

  if (!username) {
    return res.status(400).json({ error: 'Provide username in query!' });
  }
  
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
  viewAllUsers,
  viewUser,
  editProfile,
  deleteUser,
  editStatus
};