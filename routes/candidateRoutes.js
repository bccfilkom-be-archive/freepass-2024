const express = require('express');
const router = express.Router();
const candidateController = require('../controllers/candidateController');
const authMiddleware = require('../middleware/authMiddleware');

router.get('/', authMiddleware.authenticateUser, candidateController.viewCandidate);

module.exports = router;
