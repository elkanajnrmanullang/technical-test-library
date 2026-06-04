package peminjaman

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"gorm.io/gorm"
)

type PeminjamanRepository struct {
	db *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) *PeminjamanRepository {
	return &PeminjamanRepository{
		db: db,
	}
}

func (r *PeminjamanRepository) GetAllPeminjaman() ([]domain.Peminjaman, error) {
	var peminjamen []domain.Peminjaman
	err := r.db.Find(&peminjamen).Error
	return peminjamen, err
}

func (r *PeminjamanRepository) GetPeminjamanByID(id string) (*domain.Peminjaman, error) {
	var peminjaman domain.Peminjaman
	err := r.db.First(&peminjaman, "id_peminjaman = ?", id).Error
	return &peminjaman, err
}

func (r *PeminjamanRepository) CreatePeminjaman(peminjaman *domain.Peminjaman) error {
	return r.db.Create(peminjaman).Error
}

func (r *PeminjamanRepository) UpdatePeminjaman(peminjaman *domain.Peminjaman) error {
	return r.db.Model(peminjaman).
		Select("id_anggota", "tgl_pinjam", "tgl_hrs_kembali", "jaminan").
		Updates(peminjaman).Error
}

func (r *PeminjamanRepository) DeletePeminjaman(id string) error {
	return r.db.Delete(&domain.Peminjaman{}, "id_peminjaman = ?", id).Error
}

func (r *PeminjamanRepository) GetDetailPeminjaman(id string) (*domain.Peminjaman, error) {
	var peminjaman domain.Peminjaman
	err := r.db.Preload("Anggota").Preload("Details").Preload("Details.BukuDetail").First(&peminjaman, "id_peminjaman = ?", id).Error
	return &peminjaman, err
}

func (r *PeminjamanRepository) CountPeminjamanByID(id string) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Peminjaman{}).Where("id_peminjaman = ?", id).Count(&count).Error
	return count, err
}
