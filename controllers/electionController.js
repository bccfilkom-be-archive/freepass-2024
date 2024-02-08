const pool = require('../config/database');
const { executeQuery } = require('../services/db');

const castVote = (req, res) => {
  const { election_id: electionId, candidate_username: candidateUsername } = req.body;

  if (!electionId || !candidateUsername) {
    return res.status(400).json({ error: 'Provide election_id and candidate_username in the request body!' });
  }

  let query = `INSERT INTO vote (user_id, election_id, candidate_id) VALUES (?, ?, (SELECT id FROM user WHERE username = ?))`;
  let values = [req.session.userId, electionId, candidateUsername];

  executeQuery(query, values)
    .then(() => {
      res.json({ message: 'Vote casted successfully' });
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    });
};

const createElection = (req, res) => {
  const { name, start_date: startDate, end_date: endDate } = req.body;

  if (!name || !startDate || !endDate) {
    return res.status(400).json({ error: 'Provide name, start_date, and end_date in the request body!' });
  } else {
    let query = `INSERT INTO election (name, start_date, end_date) VALUES (?, ?, ?)`;
    let values = [name, startDate, endDate];

    executeQuery(query, values)
      .then(() => {
        res.json({ message: 'Election created successfully' });
      })
      .catch((error) => {
        console.error(error);
        return res.status(500).json({ error: 'Internal Server Error' });
      });
  }
};

const editElection = (req, res) => {
  const { id } = req.params;
  const { name, start_date: startDate, end_date: endDate } = req.body;

  if (!name || !startDate || !endDate) {
    return res.status(400).json({ error: 'Provide name, start_date, and end_date in the request body!' });
  }

  let query = `UPDATE election SET name = ?, start_date = ?, end_date = ? WHERE id = ?`;
  let values = [name, startDate, endDate, id];

  executeQuery(query, values)
    .then((results) => {
      if (results.affectedRows === 0) {
        return res.status(404).json({ error: 'Election not found!' });
      } else {
        res.json({ message: 'Election updated successfully' });
      }
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    });
};

const viewElection = (req, res) => {
  const { id: electionId } = req.query;

  const fetchElectionInfo = () => {
    const query = electionId ? 'SELECT * FROM election WHERE id = ?' : 'SELECT * FROM election';
    return executeQuery(query, [electionId]);
  };

  const fetchVoteCounts = () => {
    if (!electionId) {
      return fetchElectionInfo().then(elections => {
        const getVoteCountsPromises = elections.map(election => {
          return executeQuery(`
            SELECT candidate_id AS id, u.username AS username, u.name AS name, COUNT(*) AS vote_count 
            FROM vote
            JOIN user u ON candidate_id = u.id
            WHERE election_id = ?
            GROUP BY candidate_id
          `, [election.id]);
        });
        return Promise.all(getVoteCountsPromises);
      });
    } else {
      return executeQuery(`
        SELECT u.username AS username, u.name AS name, candidate_id AS id, COUNT(*) AS vote_count 
        FROM vote
        JOIN user u ON candidate_id = u.id
        WHERE election_id = ?
        GROUP BY candidate_id
      `, [electionId]).then(counts => [counts]);
    }
  };

  Promise.all([fetchElectionInfo(), fetchVoteCounts()])
    .then(([elections, counts]) => {
      if (electionId && elections.length === 0) {
        res.status(400).json({ error: 'No election found!' });
      } else {
        const result = elections.map((election, index) => ({
          id: election.id,
          name: election.name,
          start_date: election.start_date,
          end_date: election.end_date,
          counts: counts[index] || []
        }));
        res.json(result);
      }
    })
    .catch((error) => {
      res.status(500).json({ error: error });
    });
};

const deleteElection = (req, res) => {
  const { id } = req.params;

  let query = `DELETE FROM election WHERE id = ?`;

  executeQuery(query, [id])
    .then((results) => {
      if (results.affectedRows === 0) {
        return res.status(404).json({ error: 'Election not found!' });
      } else {
        res.json({ message: 'Election deleted successfully' });
      }
    })
    .catch((error) => {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    });
}

module.exports = {
  castVote,
  createElection,
  editElection,
  viewElection,
  deleteElection
}
