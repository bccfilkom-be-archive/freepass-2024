const express = require('express');
const router = express.Router();

const { getAllUser, destroyUser, getMyUser } = require('../../controllers/user/userController.js');
const { adminOnly } = require('../../middleware/AuthUser.js');
const { verifyToken } = require('../../middleware/verifyToken.js');



router.get('/', verifyToken, adminOnly, getAllUser);


router.get('/me', verifyToken, getMyUser);

// Soal No 11
// Admin can delete the user/candidate
router.delete('/admin/:id', verifyToken, adminOnly, destroyUser);

module.exports = router;