const { role } = require('../../models');


// fungsi membuat role
exports.createRole = async (req, res) => {
    
    try {
      role.create({
        name: req.body.name,
        
      });
      return res.status(201).json({ message: "Role Berhasil Dibuat" });
    } catch (error) {
      return res.status(400).json({ 
        message: error.message,
        error: error.errors.map(err => err.message) 
    });
    }
  };