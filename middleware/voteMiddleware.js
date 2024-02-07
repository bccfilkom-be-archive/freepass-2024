const pool = require('../config/database');

const checkCandidate = (req, res, next) => {
  const { candidate_username: username } = req.body;

  pool.query(`SELECT * FROM user WHERE username = ? AND status = 'candidate'`, [username], (error, results) => {
    if (error) {
      console.log(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

    if (results.length === 0) {
      return res.status(404).json({ error: 'Candidate not found!' });
    }

    next();
  });
};

const checkElection = (req, res, next) => {
  const { election_id: id } = req.body;

  pool.query(`SELECT * FROM election WHERE id = ?`, [id], (error, results) => {
    if (error) {
      console.log(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }

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
  });
};

const checkVote = (req, res, next) => {
  const { election_id: electionId } = req.body;

  pool.query(
    `SELECT * FROM vote WHERE user_id = ? AND election_id = ?`,
    [req.session.userId, electionId],
    (error, results) => {
      if (error) {
        console.log(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      }

      if (results.length === 0) {
        next();
      } else {
        return res.status(409).json({ error: `You've already voted in this election` });
      }
    }
  );
};

module.exports = {
  checkCandidate,
  checkElection,
  checkVote
}
