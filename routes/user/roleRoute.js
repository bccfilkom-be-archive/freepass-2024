const express = require('express');
const router = express.Router();

const { createRole } = require('../../controllers/user/roleController.js');


// create role tidak dibutuhkan karena langsung fungsi inisiasi melalui model/index.js
// router.post('/', createRole);



module.exports = router;