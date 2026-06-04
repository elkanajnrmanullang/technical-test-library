package auth

import (
	"errors"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/ports"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/http/v1/user/user_response"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
)

type authService struct {
	repository ports.IAuthRepository
	hash       utility.Hash
	jwtUtility *utility.JWTUtility
}

func NewAuthService(repo ports.IAuthRepository, jwtUtil *utility.JWTUtility) ports.IAuthService {
	return &authService{
		repository: repo,
		hash:       utility.Hash{},
		jwtUtility: jwtUtil,
	}
}

func (s *authService) Login(username string, password string) (*user_response.LoginResponse, error) {
	data, err := s.repository.GetDataByUsername(username)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	match, err := s.hash.VerifikasiPassword(password, data.Password)
	if err != nil || !match {
		return nil, errors.New("invalid credentials")
	}

	token, err := s.jwtUtility.GenerateToken(data.IDPegawai)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.jwtUtility.GenerateRefreshToken(data.IDPegawai)
	if err != nil {
		return nil, err
	}

	return &user_response.LoginResponse{
		Username:     data.Username,
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
