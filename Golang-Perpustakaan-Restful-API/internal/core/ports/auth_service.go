package ports

import (
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user/user_response"
	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

type IAuthService interface {
	Login(username string, password string) (*user_response.LoginResponse, error)
}

type IAuthRepository interface {
	GetDataByUsername(username string) (*domain.Pegawai, error)
}
