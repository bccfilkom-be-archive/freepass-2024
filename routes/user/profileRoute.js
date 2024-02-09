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

// UPLOAD PRODUCTS dengan filter
// const uploadProducts = multer({
//     storage: storageProducts,
//     fileFilter: imageFilter
// }).single('image');

// router.post('/', (req, res, next) => {
//     uploadProducts(req, res, function (err) {
//         if (err instanceof multer.MulterError) {
//             return res.status(400).json({
//                 message: "Error during file upload",
//                 error: err.message
//             });
//         } else if (err) {
//             return res.status(500).json({
//                 message: "Internal server error",
//                 error: err.message
//             });
//         }
//         // Jika tidak ada kesalahan, lanjut ke controller
//         next();
//     });
// }, addProduct);


const uploadProducts = multer({
    storage: storageProducts,
    fileFilter: imageFilter
});



// User can edit their profile account
router.post('/', verifyToken, uploadProducts.single('image'), updateOrCreateProfile);

module.exports = router;