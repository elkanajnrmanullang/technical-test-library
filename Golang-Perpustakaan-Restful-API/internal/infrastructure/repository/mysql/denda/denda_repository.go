package denda

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"gorm.io/gorm"
)

type DendaRepository struct {
	db *gorm.DB
}

func NewDendaRepository(db *gorm.DB) *DendaRepository {
	return &DendaRepository{
		db: db,
	}
}

func (r *DendaRepository) GetAllDenda() ([]domain.Denda, error) {
	var dendas []domain.Denda
	err := r.db.Find(&dendas).Error
	return dendas, err
}

func (r *DendaRepository) GetDendaByID(id string) (*domain.Denda, error) {
	var denda domain.Denda
	err := r.db.First(&denda, "id_denda = ?", id).Error
	return &denda, err
}

func (r *DendaRepository) CreateDenda(denda *domain.Denda) error {
	return r.db.Create(denda).Error
}

func (r *DendaRepository) UpdateDenda(denda *domain.Denda) error {
	return r.db.Model(denda).
		Select("jumlah_denda", "tglpinjam", "tglhrskembali", "tglkembali", "id_peminjaman", "id_anggota").
		Updates(denda).Error
}

func (r *DendaRepository) DeleteDenda(id string) error {
	return r.db.Delete(&domain.Denda{}, "id_denda = ?", id).Error
}

func (r *DendaRepository) CountDendaByID(id string) (int64, error) {
	var count int64
	err := r.db.Model(&domain.Denda{}).Where("id_denda = ?", id).Count(&count).Error
	return count, err
}
