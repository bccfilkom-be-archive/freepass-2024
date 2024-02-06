const express = require('express');
const router = express.Router();
const adminController = require('../controllers/adminController');
const authController = require('../controllers/authController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/logout', authMiddleware.authenticateUser, authController.logout);
router.post('/register', authController.register);
router.post('/login', authController.login);
router.put('/edit-type', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), adminController.editStatus);

module.exports = router;
