
const { Sequelize } = require("sequelize")
const db = require("../config/Database.js")

const User = require("./user.js");
const Posting = require("./posting.js");
const Review = require("./review.js");
// const Product = require("./product.js");
const {DataTypes} = Sequelize;


const Comment = {
    id:{
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        primaryKey: true,  
    },
    
    comment:{
        type: DataTypes.TEXT,
        allowNull: false
    },
    reviewId:{
        type: DataTypes.UUID,
        unique: false,
        allowNull: false,
        references: {
            model: Review,
            key: 'id'
        }
    },
}

module.exports = Comment;