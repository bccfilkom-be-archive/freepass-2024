const bcrypt   = require('bcryptjs');
const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const viewUser = (req, res) => {
  const { id, username } = req.query;
  let query = `SELECT id, nim, username, name, major, faculty, status, description FROM user WHERE 1`;
  let params = [];

  if (id) {
    query += ` AND id = ?`;
    params.push(id);
  }

  if (username) {
    query += ` AND username = ?`;
    params.push(username);
  }

  executeQuery(query, params)
    .then((results) => {
      if (results.length === 0) {
        return res.status(404).json({ error: 'User(s) not found!' });
      } else {
        return res.json(results);
      }
    })
    .catch((error) => {
      console.log(error);
      return res.status(500).json({ error: 'Internal Server Error' });
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

  updateValues.push(req.user.userId);

  const updateQuery = `UPDATE user SET ` + Object.keys(updateFields).map(key => `${key} = ?`).join(`, `) + ` WHERE id = ?`;

  executeQuery(updateQuery, updateValues)
    .then(() => {
      return res.json({ message: 'User profile updated successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
};

const deleteUser = (req, res) => {
  const { username } = req.params;

  if (!username) {
    return res.status(400).json({ error: 'Provide username!' });
  }

  if (username == req.user.username) {
    return res.status(403).json({ error: 'You cannot delete your own account' });
  }

  let query = 'DELETE FROM user WHERE username = ? ';

  executeQuery(query, [username])
    .then((results) => {
      if (results.affectedRows === 0) {
        return res.status(404).json({ error: 'User not found!' });
      }
      return res.json({ message: 'User deleted successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    });
};

const editStatus = (req, res) => {
  const { username } = req.params;
  const { status } = req.body;

  if (!username || !status) {
    return res.status(400).json({ error: 'Provide username and status!' });
  }

  if (status != 'user' && status != 'admin' && status != 'candidate') {
    return res.status(422).json({ error: 'Provide a valid status!' });
  }

  let query = `UPDATE user SET status = ? WHERE username = ?`;

  executeQuery(query, [status, username])
    .then(() => {
      return res.json({ message: 'User status updated successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    });
};

module.exports = {
  viewUser,
  editProfile,
  deleteUser,
  editStatus
};
