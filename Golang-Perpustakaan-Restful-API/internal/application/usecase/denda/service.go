package denda

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/oklog/ulid/v2"
)

type dendaService struct {
	repo ports.IDendaRepository
}

func NewDendaService(repo ports.IDendaRepository) ports.IDendaService {
	return &dendaService{
		repo: repo,
	}
}

func (s *dendaService) GetAllDenda() ([]domain.Denda, error) {
	return s.repo.GetAllDenda()
}

func (s *dendaService) GetDendaByID(id string) (*domain.Denda, error) {
	return s.repo.GetDendaByID(id)
}

func (s *dendaService) CreateDenda(denda *domain.Denda) error {
	denda.IDDenda = ulid.Make().String()
	return s.repo.CreateDenda(denda)
}

func (s *dendaService) UpdateDenda(denda *domain.Denda) error {
	count, err := s.repo.CountDendaByID(denda.IDDenda)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("denda not found")
	}
	return s.repo.UpdateDenda(denda)
}

func (s *dendaService) DeleteDenda(id string) error {
	count, err := s.repo.CountDendaByID(id)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("denda not found")
	}
	return s.repo.DeleteDenda(id)
}
