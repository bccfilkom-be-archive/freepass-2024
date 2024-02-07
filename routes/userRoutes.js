const express = require('express');
const router = express.Router();
const userController = require('../controllers/userController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/all', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userController.viewAllUsers);
router.get('/', authMiddleware.authenticateUser, userController.viewUser);
router.delete('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.deleteUser);
router.put('/', authMiddleware.authenticateUser, authMiddleware.checkUsername, userController.editProfile);
router.put('/status', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.editStatus);

module.exports = router;
