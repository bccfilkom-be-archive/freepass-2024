const jwt = require('jsonwebtoken');
const { role } = require('../models');
const {secret} = require('../config/auth.js')



exports.adminOnly = async (req, res, next) => {


    try {
        
        const Role = await role.findByPk(req.user.roleId);

        if (!Role) {
            return res.status(400).json({
                message: 'Role tidak ada yang cocok',
            });
        }

        if (Role.name !== "admin") {
            return res.status(403).json({
                message: 'Akses terlarang, Anda bukan admin',
            });
        }

        next();
    } catch (error) {
        return res.status(401).json({
            message: 'Token tidak valid',
        });
    }
};


exports.candidateOnly = async (req, res, next) => {


    try {
        
        const Role = await role.findByPk(req.user.roleId);

        if (!Role) {
            return res.status(400).json({
                message: 'Role tidak ada yang cocok',
            });
        }

        if (Role.name !== "candidate") {
            return res.status(403).json({
                message: 'Akses terlarang, Anda bukan candidate',
            });
        }

        next();
    } catch (error) {
        return res.status(401).json({
            message: 'Token tidak valid',
        });
    }
};



exports.userOnly = async (req, res, next) => {


    try {
        
        const Role = await role.findByPk(req.user.roleId);

        if (!Role) {
            return res.status(400).json({
                message: 'Role tidak ada yang cocok',
            });
        }

        if (Role.name !== "user") {
            return res.status(403).json({
                message: 'Akses terlarang, Anda bukan user',
            });
        }

        next();
    } catch (error) {
        return res.status(401).json({
            message: 'Token tidak valid',
        });
    }
};
