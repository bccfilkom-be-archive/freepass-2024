
const Role = require('./role.js');
const User = require('./user.js');
const Posting = require('./posting.js')
const Voting = require('./voting.js')
const Profile = require('./profile.js')
const Review = require('./review.js')
const Time = require('./time.js')

const db = require('../config/Database.js');
const bcrypt = require('bcryptjs');





const role = db.define("Role", Role, {
  tableName: "roles",
  underscored: true,
}
);

const user = db.define("User", User, {
  tableName: "users",
  underscored: true,
}
);


role.hasMany(user,  {
  onDelete: 'CASCADE', 

});
user.belongsTo(role, { 
  foreignKey: 'roleId' 
});



const posting = db.define("Posting", Posting, {
  tableName: "postings",
  underscored: true,
}
);


user.hasMany(posting,  {
  onDelete: 'CASCADE', 
 
});
posting.belongsTo(user, { 
  foreignKey: 'userId' 
});



const profile = db.define("Profile", Profile, {
  tableName: "profiles",
  underscored: true,
}
);


user.hasOne(profile,  {
  onDelete: 'CASCADE', // Ketika role dihapus, semua user yang terkait juga akan dihapus
  // foreignKey: {
  //   allowNull: false
  // }
});
profile.belongsTo(user, { foreignKey: 'userId' });


const review = db.define("Review", Review, {
  tableName: "reviews",
  underscored: true,
}
);



// MANY TO MANY
user.belongsToMany(posting, { through: review });
posting.belongsToMany(user, { through: review });



const voting = db.define("Voting", Voting, {
  tableName: "votings",
  underscored: true,
}
);

user.hasOne(voting,  {
  onDelete: 'CASCADE', 

});
voting.belongsTo(user, {foreignKey: 'userId'});



const time = db.define("Time", Time, {
  tableName: "times",
  underscored: true,
}
);

user.hasOne(time,  {
  onDelete: 'CASCADE', 

});
time.belongsTo(user, { foreignKey: 'userId' });


async function initial() {
  await role.create({
    name: "candidate"
  })

  await role.create({
    name: "admin"
  })

  await role.create({
    name: "user"
  })

  const userRole1 = await role.findOne({ where: { name: "admin" } });
  await user.create({
    name: "admin",
    email: "admin@gmail.com",
    ktp: "1234",
    password: bcrypt.hashSync('123456', 8),
    age: 21,
    address: "Solo",
    roleId: userRole1.id,

  })

  const userRole2 = await role.findOne({ where: { name: "user" } });
  const user1 = await user.create({
    name: "ridha ilham",
    email: "ridha@gmail.com",
    ktp: "12345",
    password: bcrypt.hashSync('123456', 8),
    roleId: userRole2.id,

  })

  const user2 = await user.create({
    name: "adi setyawan",
    email: "adi@gmail.com",
    ktp: "123456",
    password: bcrypt.hashSync('123456', 8),
    roleId: userRole2.id,

  })

  console.log(user1.id);

  await profile.create({
    age: 21,
    address: "Solo",
    image: '/assets/images/alam1.jpg',
    userId: user1.id
  })

  
  

  await profile.create({
    age: 21,
    address: "Solo",
    image: '/assets/images/alam2.jpg',
    userId: user2.id
  })

}


db.sync()
  .then(() => {
    // initial();

    console.log("database connected");

  }).catch(() => {
    console.log("database failed");
  })

module.exports = { db, user, role, posting, profile, review, voting, time };