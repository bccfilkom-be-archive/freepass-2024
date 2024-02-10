const express = require("express")
const cors = require("cors")
const session = require("express-session")
const dotenv = require("dotenv")
const db = require("./models/index.js")
const SequelizeStore = require("connect-session-sequelize")
const cookieParser = require('cookie-parser')
const bcrypt = require('bcryptjs')
const morgan = require('morgan');


const UserRoute = require('./routes/user/userRoute.js');
const AuthRoute = require('./routes/auth/AuthRoute.js');

const RoleRoute = require('./routes/user/roleRoute.js');
const PostingRoute = require('./routes/election/postingRoute.js');
const ProfileRoute = require('./routes/user/profileRoute.js');
const ReviewRoute = require('./routes/election/reviewRoute.js');
const CandidateRoute = require('./routes/user/candidateRoute.js');
const VotingRoute = require('./routes/election/votingRoute.js');
const timeRoute = require('./routes/election/timeRoute.js');

const { secret } = require('./config/auth.js')

dotenv.config();

const app = express();

const port = 3000;


app.use(cookieParser())
app.use(express.json());


app.use(morgan("dev"));
app.use(express.urlencoded({ extended: true }))


app.use('/assets', express.static('assets'));


app.use(express.json());

app.use('/api/v1/vote', VotingRoute);
app.use('/api/v1/auth', AuthRoute);
app.use('/api/v1/role', RoleRoute);
app.use('/api/v1/user', UserRoute);
app.use('/api/v1/post', PostingRoute);
app.use('/api/v1/profile', ProfileRoute);
app.use('/api/v1/review', ReviewRoute);
app.use('/api/v1/candidate', CandidateRoute);
app.use('/api/v1/time', timeRoute);



// app.use(cors());

// app.use(cors(
//     {
//     credentials: true,
//     origin: ''
// }
// ));




const APP_PORT = process.env.APP_PORT

app.listen(port, ()=> {
    console.log(`server berjalan di port ${port}`)
});
