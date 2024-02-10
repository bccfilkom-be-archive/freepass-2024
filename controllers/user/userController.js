const { user, profile, posting, review, voting, role } = require('../../models');
const fs = require('fs');

// menampilkan semua data tabel user
exports.getAllUser = async (req, res) => {
    try {
        const User = await user.findAll();

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

exports.destroyUser = async (req, res) => {

    try {
        const userId = req.params.id;

        const userData = await user.findByPk(userId);

        if(!userData) {
            return res.status(404).json({
                message: "user data tidak ditemukan"
            })
        }

        const Role = await role.findByPk(userData.roleId);

        if (!Role) {
            return res.status(400).json({
                message: 'Role tidak ada yang cocok',
            });
        }

        let infoRole;

        if (Role.name == "admin") {
            return res.status(403).json({
                message: 'Admin jangan dihapus, nanti sistem rusak',
            });
        } else if (Role.name == "candidate") {

            infoRole = "candidate"

        } else if (Role.name == "user") {

            infoRole = "user"
            
        }

        const Profile = await profile.findOne({
            where: {
                userId: userData.id
            }
        })

        
        if (!userData) {
            return res.status(404).json({
                message: "User tidak ditemukan"
            });
        }
        
        
        if (!Profile) {
            return res.status(404).json({
                message: "data profile tidak ditemukan"
            })
        } 

        const checkCandidateOnVote = voting.findOne({
            where: {
                candidate_selected: userData.candidate
            }
        })

        if(checkCandidateOnVote) {
            voting.destroy({
                where: {
                    candidate_selected: userData.candidate
                }
            })
        }

        await user.destroy({
            where: {
                id: userData.id,
                
            },
            
        });

        
        if (Profile.image) {
            fs.unlinkSync(Profile.image);
        }  
              

        return res.status(200).json({
            status: "success",
            message: infoRole + " berhasil dihapus"
        });

    } catch (error) {

        
        return res.status(500).json({
            message: "Terjadi kesalahan saat menghapus postingan",
            error: error.message
        });
    }
};



// fungsi menampilkan data tabel user yang sudah login 
exports.getMyUser = async (req, res) => {
    const currentUser = await user.findOne({
      where: {
        id: req.user.id
      },
      include: [
        {
          model: profile,
          attributes: { exclude: ["createdAt", "updatedAt", "userId"] }
        }
      ],
      attributes: { exclude: ["createdAt", "updatedAt", "password"] }
    });
  
    if (currentUser) {
      return res.status(200).json({
        data: currentUser
      })
  
    }
  
    return res.status(404).json({
      message: "User tidak ditemukan"
    })
  }