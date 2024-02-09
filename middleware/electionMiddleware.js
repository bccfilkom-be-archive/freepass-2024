const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const checkCandidate = (req, res, next) => {
  const { candidate_username: username } = req.body;
  if (!username) {
    next();
  } else {
    executeQuery(`SELECT * FROM user WHERE username = ? AND status = 'candidate'`, [username])
      .then((results) => {
        if (results.length === 0) {
          return res.status(404).json({ error: 'Candidate not found!' });
        } else {
          next();
        }
      })
      .catch((error) => {
        console.log(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      });
  }
};

const checkElection = (req, res, next) => {
  const { election_id: id, } = req.body;
  if (!id) {
    next();
  } else {
    executeQuery(`SELECT * FROM election WHERE id = ?`, [id])
      .then((results) => {
        if (results.length === 0) {
          return res.status(404).json({ error: 'Election not found!' });
        }

        const election = results[0];
        const currentDate = new Date();

        if (currentDate < new Date(election.start_date)) {
          return res.status(401).json({ error: 'Election has not started yet!' });
        }

        if (currentDate > new Date(election.end_date)) {
          return res.status(401).json({ error: 'Election has finished!' });
        }

        next();
      })
      .catch((error) => {
        console.log(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      });
  }
};

const checkVote = (req, res, next) => {
  const { election_id: electionId } = req.body;
  executeQuery(`SELECT * FROM vote WHERE user_id = ? AND election_id = ?`,
    [req.user.userId, electionId])
    .then((results) => {
      if (results.length === 0) {
        next();
      } else {
        return res.status(409).json({ error: `You've already voted in this election` });
      }
    })
    .catch((error) => {
      console.log(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    })
};

module.exports = {
  checkCandidate,
  checkElection,
  checkVote
}
