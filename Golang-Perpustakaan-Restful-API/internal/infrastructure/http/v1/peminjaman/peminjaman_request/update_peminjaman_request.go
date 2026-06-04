package peminjaman_request

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type UpdatePeminjamanRequest struct {
	IDPeminjaman string    `json:"id_peminjaman" validate:"required"`
	IDAnggota    string    `json:"id_anggota" validate:"required"`
	TglPinjam    time.Time `json:"tgl_pinjam" validate:"required"`
	TglHrsKembali time.Time `json:"tgl_hrs_kembali" validate:"required"`
	Jaminan      string    `json:"jaminan" validate:"required"`
}

func (req *UpdatePeminjamanRequest) ToDomain() *domain.Peminjaman {
	return &domain.Peminjaman{
		IDPeminjaman: req.IDPeminjaman,
		IDAnggota:    req.IDAnggota,
		TglPinjam:    req.TglPinjam,
		TglHrsKembali: req.TglHrsKembali,
		Jaminan:      req.Jaminan,
	}
}