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

const editElection = (req, res) => {
  const { id } = req.query;
  const { name, start_date: startDate, end_date: endDate } = req.body;
  pool.query(`UPDATE election SET name = ?, start_date = ?, end_date = ? WHERE id = ?`, [name, startDate, endDate, id], (error, results) => {
    if (error) {
      console.error(error);
      return res.status(500).json({ error: 'Internal Server Error' });
    }
    res.json({ message: 'Election edited successfully' });
  });
};

const viewElection = (req, res) => {
  const { id: electionId } = req.query;

  const fetchElectionInfo = () => {
    return new Promise((resolve, reject) => {
      const query = electionId ? 'SELECT * FROM election WHERE id = ?' : 'SELECT * FROM election';
      pool.query(query, [electionId], (error, elections) => {
        if (error) {
          console.error(error);
          reject('Internal Server Error');
        } else {
          resolve(elections);
        }
      });
    });
  };

  const fetchVoteCounts = () => {
    return new Promise((resolve, reject) => {
      if (!electionId) {
        resolve([]);
      } else {
        pool.query(`
          SELECT u.username AS username, u.name AS name, candidate_id AS id, COUNT(*) AS vote_count 
          FROM vote
          JOIN user u ON candidate_id = u.id
          WHERE election_id = ?
          GROUP BY candidate_id
        `, [electionId], (error, counts) => {
          if (error) {
            console.error(error);
            reject('Internal Server Error');
          } else {
            resolve(counts);
          }
        });
      }
    });
  };

  Promise.all([fetchElectionInfo(), fetchVoteCounts()])
    .then(([elections, counts]) => {
      if (electionId && elections.length === 0) {
        res.status(400).json({ error: 'No election found' });
      } else {
        res.json({ information: elections[0], counts: counts });
      }
    })
    .catch((error) => {
      res.status(500).json({ error: error });
    });
};

module.exports = {
  castVote,
  createElection,
  editElection,
  viewElection
}