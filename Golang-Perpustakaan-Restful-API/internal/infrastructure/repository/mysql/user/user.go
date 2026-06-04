package user

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (q *UserRepository) GetDataByUsername(username string) (*domain.Pegawai, error) {
	var data domain.Pegawai
	if err := q.db.Where("username = ?", username).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

