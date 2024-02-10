const express = require('express');
const router = express.Router();
const { addPosting, readPostings, detailPosting, updatePosting, destroyPosting, destroyPostingByCandidate, destroyPostingByAdmin } = require('../../controllers/election/postingController.js');

const multer = require('multer');
// const mulParse = multer();

const path = require('path');
const { verifyToken } = require('../../middleware/verifyToken.js');
const { adminOnly, candidateOnly, userOnly } = require('../../middleware/AuthUser.js');



const storageProducts = multer.diskStorage({
    destination: function (req, file, cb) {
        cb(null, "./assets/images");
    },
    filename: function (req, file, cb) {
        cb(null, file.originalname);
    }
    
});

// Filter untuk menerima hanya data gambar
const imageFilter = function (req, file, cb) {
    const allowedExtensions = ['.jpg', '.jpeg', '.png', '.gif', '.svg'];
    const fileExtension = path.extname(file.originalname).toLowerCase();
    
    if (allowedExtensions.includes(fileExtension)) {
        cb(null, true);
    } else {
        cb(new Error('File format not supported'));
    }
};



const uploadImage = multer({
    storage: storageProducts,
    fileFilter: imageFilter
});

// Soal No 14
// Candidate can create a post
router.post('/', verifyToken, candidateOnly, uploadImage.single('image'), addPosting);


// Soal No 6
// User can view the candidate's posts
router.get('/', verifyToken, userOnly, readPostings);




// Soal No 9
// Admin can view the candidate’s posts  
router.get('/admin/', verifyToken, adminOnly, readPostings);



// Soal No 4
// User
// System will display the selected candidate's post along with its details
router.get('/:id', verifyToken, userOnly, detailPosting)






// Soal No 15
// Candidate can update a post
router.put('/:id', verifyToken, candidateOnly, uploadImage.single('image'), updatePosting);


// Soal No 16
// Candidate can delete a post
router.delete('/:id', verifyToken, candidateOnly, destroyPostingByCandidate)

// Soal No 12
// Admin can delete the candidate's posts
router.delete('/admin/:id', verifyToken, adminOnly, destroyPostingByAdmin)

module.exports = router;