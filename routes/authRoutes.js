const express = require('express');
const router = express.Router();
const authController = require('../controllers/authController');
const authMiddleware = require('../middleware/authMiddleware');

router.get('/logout', authMiddleware.authenticateUser, authController.logout);
router.post('/register', authMiddleware.checkUsername, authController.register);
router.post('/login', authController.login);

module.exports = router;
