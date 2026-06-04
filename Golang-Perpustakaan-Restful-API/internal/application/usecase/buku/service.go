package buku

import (
	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
)

type BukuService struct {
	Repository ports.IBukuRepository
}

func NewBukuService(Repository ports.IBukuRepository) *BukuService {
	return &BukuService{
		Repository,
	}
}

func (s *BukuService) GetAllBuku() ([]domain.Buku, error) {
	return s.Repository.GetAllBuku()
}

func (s *BukuService) GetBukuByID(id string) (*domain.Buku, error) {
	return s.Repository.GetBukuByID(id)
}

