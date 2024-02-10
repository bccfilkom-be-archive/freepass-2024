const pool = require('../config/database');

const executeQuery = (query, values) => {
  return new Promise((resolve, reject) => {
    pool.query(query, values, (error, results) => {
      if (error) {
        console.error(error);
        reject('Internal Server Error');
      } else {
        resolve(results);
      }
    });
  });
};

module.exports = {
  executeQuery,
};
