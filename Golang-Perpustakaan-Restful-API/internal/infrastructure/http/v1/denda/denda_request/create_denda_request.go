package denda_request

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/oklog/ulid/v2"
)

type CreateDendaRequest struct {
	JumlahDenda      uint64    `json:"jumlah_denda" validate:"required"`
	Tglpinjam        time.Time `json:"tgl_pinjam"`
	Tglhrskembali    time.Time `json:"tgl_hrs_kembali"`
	Tglkembali       time.Time `json:"tgl_kembali"`
	IDPeminjaman     string    `json:"id_peminjaman" validate:"required"`
	IDAnggota        string    `json:"id_anggota" validate:"required"`
}

func (req *CreateDendaRequest) ToDomain() *domain.Denda {
	return &domain.Denda{
		IDDenda:          ulid.Make().String(),
		JumlahDenda:      req.JumlahDenda,
		Tglpinjam:        req.Tglpinjam,
		Tglhrskembali:    req.Tglhrskembali,
		Tglkembali:       req.Tglkembali,
		IDPeminjaman:     req.IDPeminjaman,
		IDAnggota:        req.IDAnggota,
	}
}
