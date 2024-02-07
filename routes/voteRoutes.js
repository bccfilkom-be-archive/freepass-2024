const express = require('express');
const router = express.Router();
const voteController = require('../controllers/voteController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');
const voteMiddleware = require('../middleware/voteMiddleware')

router.post('/', authMiddleware.authenticateUser, voteMiddleware.checkVote, voteMiddleware.checkElection, voteMiddleware.checkCandidate, voteController.castVote);

router.post('/create', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), voteController.createElection);

module.exports = router;
