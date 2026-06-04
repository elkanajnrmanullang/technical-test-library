package peminjaman_request

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/oklog/ulid/v2"
)

type CreatePeminjamanRequest struct {
	IDAnggota    string    `json:"id_anggota" validate:"required"`
	TglPinjam    time.Time `json:"tgl_pinjam" validate:"required"`
	TglHrsKembali time.Time `json:"tgl_hrs_kembali" validate:"required"`
	Jaminan      string    `json:"jaminan" validate:"required"`
}

func (req *CreatePeminjamanRequest) ToDomain() *domain.Peminjaman {
	return &domain.Peminjaman{
		IDPeminjaman: ulid.Make().String(),
		IDAnggota:    req.IDAnggota,
		TglPinjam:    req.TglPinjam,
		TglHrsKembali: req.TglHrsKembali,
		Jaminan:      req.Jaminan,
	}
}