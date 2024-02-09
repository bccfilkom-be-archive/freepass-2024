const { Op } = require('sequelize');
const { posting, review, user } = require('../../models');
const fs = require('fs');
const { profile } = require('console');
const path = require('path')
const Review = require('../../models/review.js')
const User = require('../../models/user.js')
const Profile = require('../../models/profile.js')



exports.addPosting= async (req, res) => {

    try {
        
        let image = req.file
        let idUser = req.user.id
        let { name, description } = req.body

        if (name == null || description == null || image == null) {
            return res.status(400).json({
                message: "data postingan tidak lengkap"
            })
        }

        



        const newPosting = await posting.create({
            name: name,
            description: description,
            image: req.file.path,
            userId: idUser
        });


        res.status(201).json({
            message: "berhasil menambah postingan",
            data: newPosting
        })

    } catch (error) {
        res.status(500).json(error.message)
    }

}



exports.readPostings = async (req, res) => {
    try {
      // Mengambil semua produk dari database
      const Postings = await posting.findAndCountAll();
      
      // Memeriksa jika daftar produk kosong
      if (Postings.length === 0) {
        return res.status(404).json({ message: 'Tidak ada produk yang ditemukan' });
      }
      
      // Mengirim respons dengan daftar produk yang ditemukan
      res.status(200).json({ Postings });
    } catch (error) {
      // Mengirim respons jika terjadi kesalahan
      console.error("Error while fetching products:", error);
      res.status(500).json({ error: 'Internal server error' });
    }
  };


exports.detailPosting = async (req, res) => {

    
    try {
        let idPosting = req.params.id;

        const postingData = await posting.findByPk(idPosting, {
            include: [
                {
                    model: user,
                    attributes: { exclude: ["email","name","ktp" ,"password", "createdAt", "updatedAt", "roleId", "RoleId"] }
                }
            ]
        });
        
        
        if (!postingData) {
            res.status(404).json({
                message: "postingan yang ada cari tidak ada"
            })
        }

        const commentData = await review.findAndCountAll({
            where: {
                postingId: postingData.id
            },
            attributes: {exclude: ["createdAt", "updatedAt", "UserId", "PostingId"]}
        })

        

        return res.status(200).json({
            post: postingData,
            comment: commentData    
        })

    } catch (error) {
        res.status(500).json({
            error: error.message
        })
    }

}




exports.updatePosting = async (req, res) => {


    try {
        const idParams = req.params.id;
        const { name, description } = req.body;

        const postingData = await posting.findByPk(idParams);

        if (!postingData) {
            return res.status(404).json({
                message: "Postingan dengan ID tidak ditemukan"
            });
        }



        // Update product data
        postingData.name = name || postingData.name;
        postingData.description = description || postingData.description;


        if (req.file) {
            postingData.image = req.file.path;
        }


        // Save the updated product
        await postingData.save();

        res.status(200).json({
            message: "Berhasil update postingan",
            data: postingData
        });
    } catch (error) {
        res.status(500).json({
            message: "Gagal mengupdate postingan",
            error: error.message
        });
    }
};



exports.destroyPosting = async (req, res) => {

    try {
        const postingId = req.params.id;

        const postingData = await posting.findByPk(postingId);
        
        console.log(req.user.id);
        console.log(postingData.userId);
        if (!postingData) {
            return res.status(404).json({
                message: "Postingan tidak ditemukan"
            });
        }

        if(postingData.userId !== req.user.id) {
            return res.status(404).json({
                message: "data postingan milik candidate lain"
            })
        }

        

        await posting.destroy({
            where: {
                id: postingId,

            }
        });

        
        if (postingData.image) {
            fs.unlinkSync(postingData.image);
        }  

        return res.status(200).json({
            message: "Postingan berhasil dihapus"
        });

    } catch (error) {

        
        return res.status(500).json({
            message: "Terjadi kesalahan saat menghapus postingan",
            error: error.message
        });
    }
};



