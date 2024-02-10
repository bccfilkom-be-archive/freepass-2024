const { review, posting, comment } = require('../../models');




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

    
    const { komentar } = req.body;

    if (komentar == null) {
        return res.status(400).json({
            message: "komentar belum diisi"
        });
    }


    try {
        const myReview = await review.findOne({
            where: {
                postingId: idPosting,
                userId: idUser
            }
        });

        if (myReview) {
            await comment.create({
                comment: komentar,
                reviewId: myReview.id
            });
            

            return res.status(200).json({
                message: "Komentar berhasil dibuat"
            });
        } else {

            
            const newReview = await review.create({
                postingId: idPosting,
                userId: idUser,
            });

            console.log(comment.comment);
            await comment.create({
                comment: komentar,
                reviewId: newReview.id
            });

            return res.status(201).json({
                message: "Komentar berhasil dibuat"
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

        const isCommentExist = await comment.findOne({
            where: {
                id: idComment
            }
        })

        if (!isCommentExist) {
            return res.status(400).json({
                message: "comment tidak ditemukan"
            })
        }

        await comment.destroy({
            where: {
                id: isCommentExist.id
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