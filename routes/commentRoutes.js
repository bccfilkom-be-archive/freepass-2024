const express = require('express');
const router = express.Router();
const commentController = require('../controllers/commentController');
const authMiddleware = require('../middleware/authMiddleware');
const postMiddleware = require('../middleware/postMiddleware');
const commentMiddleware = require('../middleware/commentMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.post('/', authMiddleware.authenticateUser, postMiddleware.checkPostExistence, commentController.addComment);
router.delete('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["admin"]), commentMiddleware.checkCommentExistence, commentController.deleteComment);

module.exports = router;
