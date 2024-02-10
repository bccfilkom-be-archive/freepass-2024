const express = require('express');
const router = express.Router();
const multer = require('multer');

const { createCandidate, getAllCandidate } = require('../../controllers/user/candidateController.js');
const { adminOnly, userOnly } = require('../../middleware/AuthUser.js');
const { verifyToken } = require('../../middleware/verifyToken.js');

const path = require('path');



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



// Users can view information about the candidates
router.get('/', verifyToken, userOnly, getAllCandidate)


// SOal No 8
// Admin can promote user to candidate
router.put('/admin/', verifyToken, adminOnly, uploadImage.single('image'), createCandidate);

module.exports = router;