
const { Sequelize } = require("sequelize")
const db = require("../config/Database.js");
const User = require("./user.js");

const {DataTypes} = Sequelize;


const Posting = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        primaryKey: true,  
    },
    name:{
        type: DataTypes.STRING,
        allowNull: false,
        unique: true
    },
    description:{
        type: DataTypes.STRING
    },
    image: {
        type: DataTypes.STRING,
        allowNull: false
    },
    userId: {
        type: DataTypes.UUID,
        allowNull: false,
        reference: {
            model: User,
            key: 'id'
        }
    }
}

module.exports = Posting;