const express = require('express');
const router = express.Router();
const multer = require('multer');

const { createCandidate, getAllCandidate } = require('../../controllers/user/candidateController.js');
const { adminOnly, userOnly } = require('../../middleware/AuthUser.js');
const { verifyToken } = require('../../middleware/verifyToken.js');




// Users can view information about the candidates
router.get('/', verifyToken, userOnly, getAllCandidate)


// SOal No 8
// Admin can promote user to candidate
router.put('/admin/', verifyToken, adminOnly, createCandidate);

module.exports = router;