const { review, posting } = require('../../models');



// fungsi memberikan komentar atau edit komentar lama
exports.createorUpdateReview = async (req, res) => {
    const idUser = req.user.id;
    const idPosting = req.params.id;

    const isPostingExist = await posting.findOne({
        where: {
            id: idPosting
        }
    })

    if (!isPostingExist) {
        return res.status(400).json({
            message: "postingan yang ingin anda komentari tidak ditemukan"
        })
    }

    const { comment } = req.body;



    try {
        const myReview = await review.findOne({
            where: {
                postingId: idPosting,
                userId: idUser
            }
        });

        if (myReview) {
            await myReview.update({

                comment: comment || myReview.comment
            });

            return res.status(200).json({
                message: "Comment berhasil diupdate"
            });
        } else {

            if (comment === null) {
                return res.status(400).json({
                    message: "Comment belum diisi dan anda juga belum melakukan review produk ini"
                });
            }

            await review.create({
                postingId: idPosting,
                userId: idUser,
                comment: comment
            });

            return res.status(201).json({
                message: "Comment berhasil dibuat"
            });
        }
    } catch (error) {
        console.error(error);
        return res.status(500).json({
            message: error.message
        });
    }
};


exports.deleteReview = async (req, res) => {

    try {
        const idComment = req.params.id;

        const isPostingExist = await review.findOne({
            where: {
                id: idComment
            }
        })

        if (!isPostingExist) {
            return res.status(400).json({
                message: "comment tidak ditemukan"
            })
        }

        await review.destroy({
            where: {
                id: idComment
            }
        })

        return res.status(200).json({
            message: "Berhasil delete comment"
        });

    } catch (error) {
        console.error(error);
        return res.status(500).json({
            message: error.message
        });
    }
}