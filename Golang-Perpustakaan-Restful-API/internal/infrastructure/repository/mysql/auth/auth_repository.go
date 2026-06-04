package auth

import (
	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db,
	}
}

func (r *AuthRepository) GetDataByUsername(username string) (*domain.Pegawai, error) {
	var data domain.Pegawai
	if err := r.db.Where("username = ?", username).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
