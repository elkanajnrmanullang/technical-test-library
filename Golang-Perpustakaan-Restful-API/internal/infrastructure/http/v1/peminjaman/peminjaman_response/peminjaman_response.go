package peminjaman_response

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type PeminjamanResponse struct {
	IDPeminjaman  string    `json:"id"`
	IDAnggota     string    `json:"id_anggota"`
	TglPinjam     time.Time `json:"tgl_pinjam"`
	TglHrsKembali time.Time `json:"tgl_hrs_kembali"`
	Jaminan       string    `json:"jaminan"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type PeminjamanDetailResponse struct {
	IDPeminjaman  string    `json:"id"`
	Anggota       AnggotaResponse `json:"anggota"`
	TglPinjam     time.Time `json:"tgl_pinjam"`
	TglHrsKembali time.Time `json:"tgl_hrs_kembali"`
	Jaminan       string    `json:"jaminan"`
	Details       []DetailPinjamResponse `json:"details"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type AnggotaResponse struct {
	IDAnggota string `json:"id_anggota"`
	Nama      string `json:"nama"`
}

type DetailPinjamResponse struct {
	IDDetailpinjam string `json:"id_detailpinjam"`
	IDBuku         string `json:"id_buku"`
	Kondisi        string `json:"kondisi"`
}


func FromDomain(domainPeminjaman *domain.Peminjaman) *PeminjamanResponse {
	return &PeminjamanResponse{
		IDPeminjaman:  domainPeminjaman.IDPeminjaman,
		IDAnggota:     domainPeminjaman.IDAnggota,
		TglPinjam:     domainPeminjaman.TglPinjam,
		TglHrsKembali: domainPeminjaman.TglHrsKembali,
		Jaminan:       domainPeminjaman.Jaminan,
		CreatedAt:     domainPeminjaman.CreatedAt,
		UpdatedAt:     domainPeminjaman.UpdatedAt,
	}
}

func FromDomainDetail(domainPeminjaman *domain.Peminjaman) *PeminjamanDetailResponse {
	var details []DetailPinjamResponse
	for _, detail := range domainPeminjaman.Details { // Corrected line
		details = append(details, DetailPinjamResponse{
			IDDetailpinjam: detail.IDDetailpinjam,
			IDBuku:         detail.IDBuku,
			Kondisi:        detail.Kondisi,
		})
	}
	return &PeminjamanDetailResponse{
		IDPeminjaman:  domainPeminjaman.IDPeminjaman,
		Anggota:       AnggotaResponse{IDAnggota: domainPeminjaman.Anggota.IDAnggota, Nama: domainPeminjaman.Anggota.Nama},
		TglPinjam:     domainPeminjaman.TglPinjam,
		TglHrsKembali: domainPeminjaman.TglHrsKembali,
		Jaminan:       domainPeminjaman.Jaminan,
		Details:       details,
		CreatedAt:     domainPeminjaman.CreatedAt,
		UpdatedAt:     domainPeminjaman.UpdatedAt,
	}
}

func FromDomainList(domainPeminjaman []domain.Peminjaman) []PeminjamanResponse {
	var responses []PeminjamanResponse
	for _, peminjaman := range domainPeminjaman {
		responses = append(responses, *FromDomain(&peminjaman))
	}
	return responses
}
