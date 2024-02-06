const mysql = require('mysql');

const pool = mysql.createPool({
  host: 'localhost',
  user: 'root',
  password: '',
  database: 'backend_bcc',
});

module.exports = pool;
