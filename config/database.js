const mysql = require('mysql');

const pool = mysql.createPool({
  host: 'db',
  user: 'root',
  password: '',
  database: 'backend_bcc_freepass_2024',
});

module.exports = pool;
