package denda_request

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type UpdateDendaRequest struct {
	IDDenda          string    `json:"id_denda" validate:"required"`
	JumlahDenda      uint64    `json:"jumlah_denda" validate:"required"`
	Tglpinjam        time.Time `json:"tgl_pinjam"`
	Tglhrskembali    time.Time `json:"tgl_hrs_kembali"`
	Tglkembali       time.Time `json:"tgl_kembali"`
	IDPeminjaman     string    `json:"id_peminjaman" validate:"required"`
	IDAnggota        string    `json:"id_anggota" validate:"required"`
}

func (req *UpdateDendaRequest) ToDomain() *domain.Denda {
	return &domain.Denda{
		IDDenda:          req.IDDenda,
		JumlahDenda:      req.JumlahDenda,
		Tglpinjam:        req.Tglpinjam,
		Tglhrskembali:    req.Tglhrskembali,
		Tglkembali:       req.Tglkembali,
		IDPeminjaman:     req.IDPeminjaman,
		IDAnggota:        req.IDAnggota,
	}
}
