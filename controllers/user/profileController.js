const { profile, user } = require('../../models');
const path = require('path')
const bcrypt = require('bcryptjs')


// membuat dan edit tabel profile
exports.updateOrCreateProfile = async (req, res) => {
    const { age, address, ktp } = req.body


    const idUser = req.user.id;
    const image = req.file.path

    const profileData = await profile.findOne({
        where: {
            userId: idUser
        }
    })

    const userData = await user.findOne({
        where: {
            id: idUser
        }
    })

    if(!userData) {
        return res.status(404).json({
            message: "user data tidak ditemukan"
        })
    }

    

    if(profileData) {
        const hashPassword = await bcrypt.hash(req.body.password, 8)
        await user.update({
            password: hashPassword || userData.password,
            ktp: ktp || userData.ktp
        }, {
            where: {
                id: idUser 
            }
        })

        await profile.update({
            age: age || profileData.age,
            address: address || profileData.address,
            image: image || profileData.image
        }, {
            where: {
                userId: idUser
            }
        })

        return res.status(200).json({
            message : "Profile berhasil update"
        })
        
    } else {
        
        await profile.create({
            age: age,
            address: address,
            userId: idUser
        })
        return res.status(200).json({
            message : "Profile berhasil dibuat"
        })
        
    }

    
}