const express = require('express');
const router = express.Router();
const commentController = require('../controllers/commentController');
const authMiddleware = require('../middleware/authMiddleware');
const postMiddleware = require('../middleware/postMiddleware');
const commentMiddleware = require('../middleware/commentMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/', authMiddleware.authenticateUser, commentController.viewComment);
router.post('/', authMiddleware.authenticateUser, authMiddleware.getUserInfo, postMiddleware.checkPostExistence, commentController.addComment);
router.delete('/:id', authMiddleware.authenticateUser, authMiddleware.getUserInfo, userMiddleware.checkUserStatus(["admin"]), commentMiddleware.checkCommentExistence, commentController.deleteComment);

module.exports = router;
