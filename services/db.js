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

const getUserInfo = async (id) => {
  try {
    const results = await executeQuery(`SELECT * FROM user WHERE id = ?`, [id]);
    return results[0];
  } catch (error) {
    console.error(error);
    throw error;
  }
};

module.exports = {
  executeQuery,
  getUserInfo
};
