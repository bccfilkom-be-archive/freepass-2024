const express = require('express');
const router = express.Router();
const voteController = require('../controllers/electionController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');
const voteMiddleware = require('../middleware/electionMiddleware')

router.post('/vote', authMiddleware.authenticateUser, authMiddleware.getUserInfo, voteMiddleware.checkVote, voteMiddleware.checkElection, voteMiddleware.checkCandidate, voteController.castVote);

router.get('/', authMiddleware.authenticateUser, voteController.viewElection);
router.post('/', authMiddleware.authenticateUser, authMiddleware.getUserInfo, userMiddleware.checkUserStatus(["admin"]), voteController.createElection);
router.put('/:id', authMiddleware.authenticateUser, authMiddleware.getUserInfo, userMiddleware.checkUserStatus(["admin"]), voteController.editElection);
router.delete('/:id', authMiddleware.authenticateUser, authMiddleware.getUserInfo, userMiddleware.checkUserStatus(["admin"]), voteController.deleteElection);

module.exports = router;
