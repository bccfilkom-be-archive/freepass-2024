const express = require('express');
const router = express.Router();
const authController = require('../controllers/authController');
const authMiddleware = require('../middleware/authMiddleware');

router.delete('/', authMiddleware.authenticateUser, authController.logout);
router.post('/', authController.login);

module.exports = router;
