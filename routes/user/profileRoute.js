const express = require('express');
const router = express.Router();
const { updateOrCreateProfile } = require('../../controllers/user/profileController.js');
const { verifyToken } = require('../../middleware/verifyToken.js');
const multer = require('multer');
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


// Soal No 3
// User can edit their profile account
router.post('/', verifyToken, uploadImage.single('image'), updateOrCreateProfile);

module.exports = router;