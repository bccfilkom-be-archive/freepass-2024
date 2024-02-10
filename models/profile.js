
const { Sequelize } = require("sequelize");
const db = require("../config/Database.js");
const User = require("./user.js");
const {DataTypes} = Sequelize;


const Profile = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        
        primaryKey: true,
        
    },
    age:{
        type: DataTypes.INTEGER,
        allowNull: false
    },

    address:{
        type: DataTypes.TEXT,
        allowNull: false,
    },
    image:{
        type: DataTypes.STRING,
        
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



module.exports = Profile;