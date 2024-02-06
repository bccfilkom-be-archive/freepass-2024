const express = require('express');
const router = express.Router();
const userController = require('../controllers/userController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/all', authMiddleware.authenticateUser, userController.viewAllUsers);
router.get('/', authMiddleware.authenticateUser, userMiddleware.checkUserExistence, userController.viewUser);
router.delete('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.deleteUser);
router.put('/', authMiddleware.authenticateUser, userController.editProfile);

module.exports = router;
