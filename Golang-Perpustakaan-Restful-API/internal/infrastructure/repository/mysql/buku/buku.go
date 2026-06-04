package buku

import (
	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"gorm.io/gorm"
)

type BukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) *BukuRepository {
	return &BukuRepository{
		db,
	}
}

func (r *BukuRepository) GetAllBuku() ([]domain.Buku, error) {
	var bukus []domain.Buku
	err := r.db.Find(&bukus).Error
	return bukus, err
}

func (r *BukuRepository) GetBukuByID(id string) (*domain.Buku, error) {
	var buku domain.Buku
	err := r.db.First(&buku, "id_buku = ?", id).Error
	return &buku, err
}
