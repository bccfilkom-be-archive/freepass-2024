
const { Sequelize } = require("sequelize")
const db = require("../config/Database.js")

const {DataTypes} = Sequelize;


const Role = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        primaryKey: true,
        
    },
    name:{
        type: DataTypes.STRING,
        allowNull: false,
        unique: true
    }
}



module.exports = Role;