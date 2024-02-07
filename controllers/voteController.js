const pool = require('../config/database');

const castVote = (req, res) => {
  const { election_id: electionId, candidate_username: candidateUsername } = req.body;
  pool.query(`INSERT INTO vote (user_id, election_id, candidate_id) VALUES (?, ?, (SELECT id FROM user WHERE username = ?))`, [req.session.userId, electionId, candidateUsername], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Vote casted successfully' });
  });
};

const createElection = (req, res) => {
  const { name, start_date: startDate, end_date: endDate } = req.body;
  pool.query(`INSERT INTO election (name, start_date, end_date) VALUES (?, ?, ?)`, [name, startDate, endDate], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Election created successfully' });
  });
};

module.exports = {
  castVote,
  createElection
}