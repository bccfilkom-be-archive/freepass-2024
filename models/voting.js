
const { Sequelize } = require("sequelize");
const db = require("../config/Database.js");
const User = require("./user.js");
const {DataTypes} = Sequelize;


const Voting = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        
        primaryKey: true,
        
    },
    candidate_selected:{
        type: DataTypes.INTEGER,
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



module.exports = Voting;