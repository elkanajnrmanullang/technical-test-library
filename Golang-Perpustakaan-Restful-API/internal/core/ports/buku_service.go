package ports

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"

type IBukuService interface {
	// jenis buku

	GetAllJenisBuku() ([]domain.Jenis_Buku, error)
	FindJenisBuku(c string) ([]domain.Jenis_Buku, error)
	GetJenisBukuById(id string) (*domain.Jenis_Buku, error)
	CreateJenisBuku(data domain.Jenis_Buku) (domain.Jenis_Buku, error)
	UpdateJenisBuku(data domain.Jenis_Buku) (domain.Jenis_Buku, error)
	HapusJenisBuku(id string) error

	// penerbit buku

	GetAllPenerbitBuku() ([]domain.Penerbit_Buku, error)
	FindPenerbitBuku(c string) ([]domain.Penerbit_Buku, error)
	GetPenerbitBukuById(id string) (*domain.Penerbit_Buku, error)
	CreatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error)
	UpdatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error)
	HapusPenerbitBuku(id string) error

	// penulis buku

	GetAllPenulisBuku() ([]domain.Penulis_Buku, error)
	FindPenulisBuku(c string) ([]domain.Penulis_Buku, error)
	GetPenulisBukuById(id string) (*domain.Penulis_Buku, error)
	CreatePenulisBuku(data domain.Penulis_Buku) (domain.Penulis_Buku, error)
	UpdatePenulisBuku(data domain.Penulis_Buku) (domain.Penulis_Buku, error)
	HapusPenulisBuku(id string) error

	// buku
	GetAllBuku() ([]domain.Buku, error)
	GetBukuByID(id string) (*domain.Buku, error)
}

//go:generate mockery --name IBukuRepository
type IBukuRepository interface {
	// jenis buku

	GetAllJenisBuku() ([]domain.Jenis_Buku, error)
	CariJenisBuku(c string) ([]domain.Jenis_Buku, error)
	GetJenisBukuById(id string) (*domain.Jenis_Buku, error)
	CreateJenisBuku(data domain.Jenis_Buku) (domain.Jenis_Buku, error)
	UpdateJenisBuku(data domain.Jenis_Buku) (domain.Jenis_Buku, error)
	DeleteJenisBuku(id string) error
	HitungDataJenisBuku(id string) int64

	// penerbit buku

	GetAllPenerbitBuku() ([]domain.Penerbit_Buku, error)
	CariPenerbitBuku(c string) ([]domain.Penerbit_Buku, error)
	GetPenerbitBukuById(id string) (*domain.Penerbit_Buku, error)
	CreatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error)
	UpdatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error)
	DeletePenerbitBuku(id string) error
	HitungDataPenerbitBuku(id string) int64

	// penulis buku

	CreatePenulisBuku(data domain.Penulis_Buku) (domain.Penulis_Buku, error)
	UpdatePenulisBuku(data domain.Penulis_Buku) (domain.Penulis_Buku, error)
	DeletePenulisBuku(id string) error
	HitungDataPenulisBuku(id string) int64
	GetAllPenulisBuku() ([]domain.Penulis_Buku, error)
	CariPenulisBuku(c string) ([]domain.Penulis_Buku, error)
	GetPenulisBukuById(id string) (*domain.Penulis_Buku, error)

	// buku
	GetAllBuku() ([]domain.Buku, error)
	GetBukuByID(id string) (*domain.Buku, error)
}
