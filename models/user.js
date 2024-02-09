
const { Sequelize } = require("sequelize");
const db = require("../config/Database.js");
const Role = require('./role.js');
const {DataTypes} = Sequelize;


const User = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        primaryKey: true,
        
    },
    name:{
        type: DataTypes.STRING,
        allowNull: false
    },
    ktp:{
        type: DataTypes.STRING, 
        unique: true,
        allowNull: false
    },
    email:{
        type: DataTypes.STRING,
        unique: true,
        allowNull: false,
        validate:{
            isEmail: true
        }
    },
    password:{
        type: DataTypes.STRING,
        allowNull: false,
        validate: {
            len: [6]
        }
    },
    candidate: {
        type: DataTypes.INTEGER,
        allowNull: true,
        unique: true,
        defaultValue: null
    },
    paslon:{
        type: DataTypes.STRING,
        unique: true,
        allowNull: true
    },
    roleId: {
        type: DataTypes.UUID,
        allowNull: false,
        references: {
            model: Role,
            key: 'id'
        }
    }
}



module.exports = User;