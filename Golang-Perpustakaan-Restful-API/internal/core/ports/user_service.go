package ports

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user/user_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type IUserService interface {
	Login(username string, password string) (*user_response.LoginResponse, error)
}

type IUserRepository interface {
	GetDataByUsername(username string) (*domain.Pegawai, error)
}
