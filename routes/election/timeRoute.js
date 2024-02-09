const express = require('express');
const router = express.Router();

const { updateOrCreateTime } = require("../../controllers/election/timeController");
const { adminOnly } = require("../../middleware/AuthUser");
const { verifyToken } = require("../../middleware/verifyToken");


// Soal No 10
// Admin can set the start and end dates for the election period
router.post('/admin/', verifyToken, adminOnly, updateOrCreateTime)

module.exports = router