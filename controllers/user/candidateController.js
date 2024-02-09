const { user, profile } = require('../../models');
const bcrypt = require('bcryptjs');
const { role } = require('../../models');




// admin mengubah status pengguna biasa menjadi status candidate
exports.createCandidate = async (req, res) => {

    

    const { name_paslon, email_paslon, candidate } = req.body



    if (name_paslon == null || email_paslon == null || candidate == null ) {
        return res.status(400).json({
            message: "mohon diisi dengan lengkap",
        });
    }

    const isUserExist = await user.findOne({
        where: {
            email: email_paslon,
        },
    });



    if (!isUserExist) {
        return res.status(400).json({
            message: "user tidak ditemukan, tidak bisa menjadi candidate",
        });
    }



    const userRole = await role.findOne({
        where: {
            name: "candidate"
        }
    });


    try {
        await user.update({

            roleId: userRole.id,
            candidate: candidate,
            paslon: name_paslon

        },
            {
                where: {
                    email: email_paslon
                }
            }
        );

        return res.status(200).json({ message: "Kandidat berhasil dibuat" });

        //   createSendToken(newUser, 201, res);

        // const token = jwt.sign(newUser.id,
        //   process.env.SECRET_KEY, 
        //   {
        //       expiresIn: '1h',
        //   }
        // );
    } catch (error) {
        return res.status(500).json({ message: error.message });
    }
};


// menampilkan semua data candidate
exports.getAllCandidate = async (req, res) => {
    try {
        const User = await user.findAll({
            where: {
                role: "candidate"
            }
        });

        return res.status(200).json({
            status: "Success",
            data: User
        })
    } catch (error) {
        return res.status(500).json({
            status: 'Fail',
            error: 'Server Down'
        })
    }

}