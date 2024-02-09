const { time } = require('../models');

exports.checkDeadline = async (req, res, next) => {
    try {
        const currentTime = new Date(); // Mendapatkan waktu saat ini

        // Mendapatkan data waktu dari database
        const deadline = await time.findOne();


        if (!deadline) {
            return res.status(400).json({
                message: "Data deadline tidak ditemukan."
            });
        }

        const startTime = new Date(deadline.start_time);
        const endTime = new Date(deadline.end_time);

        // Memeriksa apakah waktu saat ini berada di antara waktu mulai dan waktu selesai
        if (currentTime < startTime || currentTime > endTime) {
            return res.status(403).json({ 
                message: 'Tidak bisa voting karena di luar waktu pemilu' 
            });
        }

        // Jika waktu saat ini berada di dalam jangka waktu yang diizinkan, lanjutkan dengan middleware berikutnya
        next();

    } catch (error) {
        // Menangani kesalahan dengan mengirimkan respon status 500 (internal server error)
        return res.status(500).json({
            message: "Terjadi kesalahan dalam mengecek deadline."
        });
    }
};
