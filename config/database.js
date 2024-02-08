const mysql = require('mysql');

const pool = mysql.createPool({
  host: 'db',
  user: 'root',
  password: '',
  database: 'backend_bcc',
});

module.exports = pool;
