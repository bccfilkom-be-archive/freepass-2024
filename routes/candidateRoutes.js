const express = require('express');
const router = express.Router();
const candidateController = require('../controllers/candidateController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/', authMiddleware.authenticateUser, userMiddleware.checkUserExistence, candidateController.viewCandidate);

module.exports = router;
