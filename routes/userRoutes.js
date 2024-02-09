const express = require('express');
const router = express.Router();
const userController = require('../controllers/userController');
const authMiddleware = require('../middleware/authMiddleware');
const userMiddleware = require('../middleware/userMiddleware');
const authController = require('../controllers/authController');

router.get('/', authMiddleware.authenticateUser, userController.viewUser);
router.delete('/:username', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.deleteUser);
router.put('/', authMiddleware.authenticateUser, authMiddleware.checkUsername, userController.editProfile);
router.put('/status/:username', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), userMiddleware.checkUserExistence, userController.editStatus);
router.post('/', authMiddleware.checkUsername, authController.register);

module.exports = router;
