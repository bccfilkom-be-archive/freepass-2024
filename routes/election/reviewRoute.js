const express = require('express');
const router = express.Router();
const { verifyToken } = require('../../middleware/verifyToken.js');
const { createorUpdateReview, deleteReview } = require('../../controllers/election/reviewController.js');
const { userOnly, adminOnly } = require('../../middleware/AuthUser.js');


// Soal No 13
// Admin delete a user comment
router.delete('/admin/:id', verifyToken, adminOnly, deleteReview);


// Soal No 5
// User can comment on candidate’s posts
router.post('/:id', verifyToken, userOnly, createorUpdateReview);


module.exports = router;