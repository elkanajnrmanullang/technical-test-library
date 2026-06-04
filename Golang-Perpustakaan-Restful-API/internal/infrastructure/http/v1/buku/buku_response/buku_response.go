package buku_response

import (
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type BukuResponse struct {
	IDBuku          string    `json:"id_buku"`
	ISBN            string    `json:"isbn"`
	IDKategoriJenis string    `json:"id_kategori_buku"`
	Judul           string    `json:"judul_buku"`
	IDPenulisBuku   string    `json:"id_penulis_buku"`
	IDPenerbitBuku  string    `json:"id_penerbit_buku"`
	ThnTerbit       string    `json:"tahun_terbit"`
	StokBuku        int32     `json:"stok_buku"`
	RakBuku         string    `json:"rak_buku"`
	DeskripsiBuku   string    `json:"deskripsi_buku"`
	Gambarbuku      string    `json:"gambar_buku"`
	Kondisibuku     string    `json:"kondisi_buku"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func FromDomainBuku(domainBuku *domain.Buku) *BukuResponse {
	return &BukuResponse{
		IDBuku:          domainBuku.IDBuku,
		ISBN:            domainBuku.ISBN,
		IDKategoriJenis: domainBuku.IDKategoriJenis,
		Judul:           domainBuku.Judul,
		IDPenulisBuku:   domainBuku.IDPenulisBuku,
		IDPenerbitBuku:  domainBuku.IDPenerbitBuku,
		ThnTerbit:       domainBuku.ThnTerbit,
		StokBuku:        domainBuku.StokBuku,
		RakBuku:         domainBuku.RakBuku,
		DeskripsiBuku:   domainBuku.DeskripsiBuku,
		Gambarbuku:      domainBuku.Gambarbuku,
		Kondisibuku:     domainBuku.Kondisibuku,
		CreatedAt:       domainBuku.CreatedAt,
		UpdatedAt:       domainBuku.UpdatedAt,
	}
}

func FromDomainListBuku(domainBuku []domain.Buku) []BukuResponse {
	var responses []BukuResponse
	for _, buku := range domainBuku {
		responses = append(responses, *FromDomainBuku(&buku))
	}
	return responses
}
