const express = require('express');
const router = express.Router();


const { updateOrCreateVoting, countVoting } = require("../../controllers/election/votingController");
const { adminOnly } = require("../../middleware/AuthUser");
const { checkDeadline } = require("../../middleware/checkDeadline");
const { verifyToken } = require("../../middleware/verifyToken");

// menghitung hasil voting berdasarkan nomor candidate
router.get('/', verifyToken, countVoting)



// Users can cast their votes for candidates during the specified election period
router.post('/', verifyToken, checkDeadline, updateOrCreateVoting)

module.exports = router


