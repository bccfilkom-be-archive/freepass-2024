const express = require('express');
const router = express.Router();
const postController = require('../controllers/postController');
const authMiddleware = require('../middleware/authMiddleware');
const postMiddleware = require('../middleware/postMiddleware');
const userMiddleware = require('../middleware/userMiddleware');

router.get('/all', authMiddleware.authenticateUser, postController.viewAllPosts);
router.get('/', authMiddleware.authenticateUser, postController.viewPost);
router.post('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["candidate"]), postController.addPost);
router.put('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["candidate"]), postMiddleware.checkPostExistence, postMiddleware.checkPostOwnership, postController.editPost);
router.delete('/', authMiddleware.authenticateUser, userMiddleware.checkUserStatus(["candidate", "admin"]), postMiddleware.checkPostExistence, postMiddleware.checkPostOwnership, postController.deletePost);

module.exports = router;
