package ports

import "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"

type IDendaService interface {
	GetAllDenda() ([]domain.Denda, error)
	GetDendaByID(id string) (*domain.Denda, error)
	CreateDenda(denda *domain.Denda) error
	UpdateDenda(denda *domain.Denda) error
	DeleteDenda(id string) error
}

type IDendaRepository interface {
	GetAllDenda() ([]domain.Denda, error)
	GetDendaByID(id string) (*domain.Denda, error)
	CreateDenda(denda *domain.Denda) error
	UpdateDenda(denda *domain.Denda) error
	DeleteDenda(id string) error
	CountDendaByID(id string) (int64, error)
}
