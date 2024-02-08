const express = require('express');
const router = express.Router();
const userController = require('../controllers/userController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userController.viewUser);
router.get('/:username', authMiddleware.authenticateUser, userController.viewUser);
router.delete('/:username', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.deleteUser);
router.put('/', authMiddleware.authenticateUser, authMiddleware.checkUsername, userController.editProfile);
router.put('/status/:username', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.editStatus);

module.exports = router;
