package ports

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"

type IPeminjamanService interface {
	GetAllPeminjaman() ([]domain.Peminjaman, error)
	GetPeminjamanByID(id string) (*domain.Peminjaman, error)
	CreatePeminjaman(peminjaman *domain.Peminjaman) error
	UpdatePeminjaman(peminjaman *domain.Peminjaman) error
	DeletePeminjaman(id string) error
	GetDetailPeminjaman(id string) (*domain.Peminjaman, error)
}

type IPeminjamanRepository interface {
	GetAllPeminjaman() ([]domain.Peminjaman, error)
	GetPeminjamanByID(id string) (*domain.Peminjaman, error)
	CreatePeminjaman(peminjaman *domain.Peminjaman) error
	UpdatePeminjaman(peminjaman *domain.Peminjaman) error
	DeletePeminjaman(id string) error
	GetDetailPeminjaman(id string) (*domain.Peminjaman, error)
	CountPeminjamanByID(id string) (int64, error)
}
