const { user, profile } = require('../../models');
const bcrypt = require('bcryptjs');
const { role } = require('../../models');
const path = require('path')



// admin mengubah status pengguna biasa menjadi status candidate
exports.createCandidate = async (req, res) => {

    

    const { name_paslon, email_paslon, candidate } = req.body
    const image = req.file


    if (name_paslon == null || email_paslon == null || candidate == null ) {
        return res.status(400).json({
            message: "mohon diisi dengan lengkap",
        });
    }

    if (!image) {
        return res.status(400).json({
          message: "image tidak ada"
        })
      }

    const checkNomorCandidate = await user.findOne({
        where: {
            candidate: candidate
        }
    })

    if(checkNomorCandidate) {
        return res.status(403).json({
            message: "nomor candidate tidak bisa dipakai karena sudah pakai candidate lain"
        })
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

    let remider;
    

    const userRole = await role.findOne({
        where: {
            name: "candidate"
        }
    });

    if(!userRole) {
        return res.status(404).json({
            message: "role tidak ditemukan"
        })
    }

    const profileData = await profile.findOne({
        where: {
            userId: isUserExist.id
        }
    })

    if(!profileData) {
        return res.status(404).json({
            message: "profile data tidak ditemukan"
        })
    }

    console.log("halooo 1 ");
    try {
        if(isUserExist.candidate) {
            await user.update({

                roleId: userRole.id,
                candidate: isUserExist.candidate,
                paslon: name_paslon
    
            },
                {
                    where: {
                        email: email_paslon
                    }
                }
            );
            remider = `nomor candidate tidak bisa diganti dan tetap ${isUserExist.candidate}`
            
        } else {
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
        }
        
        
        
        await profile.create({
            age: profileData.age,
            address: profileData.address,
            image: image.path,
            userId: isUserExist.id
          })

        return res.status(200).json({ 
            message: "Kandidat berhasil dibuat",
            remider: remider 
        });

        
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

