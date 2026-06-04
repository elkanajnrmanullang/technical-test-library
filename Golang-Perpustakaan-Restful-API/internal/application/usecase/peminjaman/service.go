package peminjaman

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/oklog/ulid/v2"
)

type peminjamanService struct {
	repo ports.IPeminjamanRepository
}

func NewPeminjamanService(repo ports.IPeminjamanRepository) ports.IPeminjamanService {
	return &peminjamanService{
		repo: repo,
	}
}

func (s *peminjamanService) GetAllPeminjaman() ([]domain.Peminjaman, error) {
	return s.repo.GetAllPeminjaman()
}

func (s *peminjamanService) GetPeminjamanByID(id string) (*domain.Peminjaman, error) {
	return s.repo.GetPeminjamanByID(id)
}

func (s *peminjamanService) CreatePeminjaman(peminjaman *domain.Peminjaman) error {
	peminjaman.IDPeminjaman = ulid.Make().String()
	return s.repo.CreatePeminjaman(peminjaman)
}

func (s *peminjamanService) UpdatePeminjaman(peminjaman *domain.Peminjaman) error {
	count, err := s.repo.CountPeminjamanByID(peminjaman.IDPeminjaman)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("peminjaman not found")
	}
	return s.repo.UpdatePeminjaman(peminjaman)
}

func (s *peminjamanService) DeletePeminjaman(id string) error {
	count, err := s.repo.CountPeminjamanByID(id)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("peminjaman not found")
	}
	return s.repo.DeletePeminjaman(id)
}

func (s *peminjamanService) GetDetailPeminjaman(id string) (*domain.Peminjaman, error) {
	return s.repo.GetDetailPeminjaman(id)
}
