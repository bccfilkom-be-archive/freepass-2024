const { Sequelize } = require("sequelize");
const db = require("../config/Database.js");
const User = require("./user.js");

const { DataTypes } = Sequelize;

const Time = {
    id: {
        type: DataTypes.UUID,
        defaultValue: DataTypes.UUIDV4,
        primaryKey: true,
    },
    start_time: {
        type: DataTypes.DATE, 
        allowNull: false,
    },
    end_time: {
        type: DataTypes.DATE, 
        allowNull: false,
    },
    userId: {
        type: DataTypes.UUID,
        allowNull: false,
        reference: {
            model: User,
            key: 'id'
        }
    }
};

module.exports = Time;
