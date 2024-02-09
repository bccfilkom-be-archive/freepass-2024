
const { Sequelize } = require("sequelize")
const db = require("../config/Database.js")

const User = require("./user.js");
const Posting = require("./posting.js");
// const Product = require("./product.js");
const {DataTypes} = Sequelize;


const Review = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        primaryKey: true,  
    },
    userId:{
        type: DataTypes.UUID,
        unique: false,
        allowNull: false,
        references: {
            model: User,
            key: 'id'
        }
    },
    postingId:{
        type: DataTypes.UUID,
        unique: false,
        allowNull: false,
        references: {
            model: Posting,
            key: 'id'
        }
    },
  

    comment:{
        type: DataTypes.TEXT,
        allowNull: false
    }
}

module.exports = Review;