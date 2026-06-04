package denda_response

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type DendaResponse struct {
	IDDenda          string    `json:"id_denda"`
	JumlahDenda      uint64    `json:"jumlah_denda"`
	Tglpinjam        time.Time `json:"tgl_pinjam"`
	Tglhrskembali    time.Time `json:"tgl_hrs_kembali"`
	Tglkembali       time.Time `json:"tgl_kembali"`
	IDPeminjaman     string    `json:"id_peminjaman"`
	IDAnggota        string    `json:"id_anggota"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

func FromDomain(domainDenda *domain.Denda) *DendaResponse {
	return &DendaResponse{
		IDDenda:          domainDenda.IDDenda,
		JumlahDenda:      domainDenda.JumlahDenda,
		Tglpinjam:        domainDenda.Tglpinjam,
		Tglhrskembali:    domainDenda.Tglhrskembali,
		Tglkembali:       domainDenda.Tglkembali,
		IDPeminjaman:     domainDenda.IDPeminjaman,
		IDAnggota:        domainDenda.IDAnggota,
		CreatedAt:        domainDenda.CreatedAt,
		UpdatedAt:        domainDenda.UpdatedAt,
	}
}

func FromDomainList(domainDenda []domain.Denda) []DendaResponse {
	var responses []DendaResponse
	for _, denda := range domainDenda {
		responses = append(responses, *FromDomain(&denda))
	}
	return responses
}
