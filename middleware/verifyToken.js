const jwt = require('jsonwebtoken');
const config = require('../config/auth.js');
const { secret } = require('../config/auth.js');
const { user } = require('../models'); // Assuming your model is named 'User' with a capital 'U'

exports.verifyToken = async (req, res, next) => {
  try {
    let token;
    token = req.cookies.jwt;

    

    if (!token) {
      return res.status(403).json({
        message: 'Silakan login terlebih dahulu'
      });
    }

    const decoded = jwt.verify(token, process.env.SECRET_KEY);

    

    const User = await user.findOne({ where: { id: decoded.id } });

    

    if (!User) {
      return res.status(404).json({
        message: 'User not found'
      });
    }
    
    
    req.user = User;
    next();
  } catch (error) {
    return res.status(403).json({
      message: 'Kredensial salah'
    });
  }
};


