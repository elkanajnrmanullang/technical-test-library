package peminjaman_test

import (
	"errors"
	"testing"
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/peminjaman"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/peminjaman/mocks"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPeminjamanService_GetAllPeminjaman(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockPeminjaman := []domain.Peminjaman{
		{
			IDPeminjaman:  "1",
			IDAnggota:     "anggota1",
			TglPinjam:     time.Now(),
			TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
			Jaminan:       "KTP",
		},
	}

	mockRepo.On("GetAllPeminjaman").Return(mockPeminjaman, nil)

	peminjamen, err := service.GetAllPeminjaman()

	assert.Nil(t, err)
	assert.Equal(t, mockPeminjaman, peminjamen)
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_GetAllPeminjaman_Error(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("GetAllPeminjaman").Return(nil, errors.New("database error"))

	peminjamen, err := service.GetAllPeminjaman()

	assert.NotNil(t, err)
	assert.Nil(t, peminjamen)
	assert.EqualError(t, err, "database error")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_GetPeminjamanByID(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockPeminjaman := &domain.Peminjaman{
		IDPeminjaman:  "1",
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "KTP",
	}

	mockRepo.On("GetPeminjamanByID", "1").Return(mockPeminjaman, nil)

	peminjaman, err := service.GetPeminjamanByID("1")

	assert.Nil(t, err)
	assert.Equal(t, mockPeminjaman, peminjaman)
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_GetPeminjamanByID_Error(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("GetPeminjamanByID", "1").Return(nil, errors.New("not found"))

	peminjaman, err := service.GetPeminjamanByID("1")

	assert.NotNil(t, err)
	assert.Nil(t, peminjaman)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_CreatePeminjaman(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	newPeminjaman := &domain.Peminjaman{
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "KTP",
	}

	mockRepo.On("CreatePeminjaman", mock.AnythingOfType("*domain.Peminjaman")).Return(nil)

	err := service.CreatePeminjaman(newPeminjaman)

	assert.Nil(t, err)
	assert.NotEmpty(t, newPeminjaman.IDPeminjaman)
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_CreatePeminjaman_Error(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	newPeminjaman := &domain.Peminjaman{
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "KTP",
	}

	mockRepo.On("CreatePeminjaman", mock.AnythingOfType("*domain.Peminjaman")).Return(errors.New("create error"))

	err := service.CreatePeminjaman(newPeminjaman)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "create error")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_UpdatePeminjaman(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	updatedPeminjaman := &domain.Peminjaman{
		IDPeminjaman:  "1",
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "SIM",
	}

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(1), nil)
	mockRepo.On("UpdatePeminjaman", updatedPeminjaman).Return(nil)

	err := service.UpdatePeminjaman(updatedPeminjaman)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_UpdatePeminjaman_NotFound(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	updatedPeminjaman := &domain.Peminjaman{
		IDPeminjaman:  "1",
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "SIM",
	}

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(0), nil)

	err := service.UpdatePeminjaman(updatedPeminjaman)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "peminjaman not found")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_UpdatePeminjaman_CountError(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	updatedPeminjaman := &domain.Peminjaman{
		IDPeminjaman:  "1",
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "SIM",
	}

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(0), errors.New("count error"))

	err := service.UpdatePeminjaman(updatedPeminjaman)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "count error")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_UpdatePeminjaman_UpdateError(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	updatedPeminjaman := &domain.Peminjaman{
		IDPeminjaman:  "1",
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "SIM",
	}

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(1), nil)
	mockRepo.On("UpdatePeminjaman", updatedPeminjaman).Return(errors.New("update error"))

	err := service.UpdatePeminjaman(updatedPeminjaman)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "update error")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_DeletePeminjaman(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(1), nil)
	mockRepo.On("DeletePeminjaman", "1").Return(nil)

	err := service.DeletePeminjaman("1")

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_DeletePeminjaman_NotFound(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(0), nil)

	err := service.DeletePeminjaman("1")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "peminjaman not found")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_DeletePeminjaman_CountError(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(0), errors.New("count error"))

	err := service.DeletePeminjaman("1")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "count error")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_DeletePeminjaman_DeleteError(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("CountPeminjamanByID", "1").Return(int64(1), nil)
	mockRepo.On("DeletePeminjaman", "1").Return(errors.New("delete error"))

	err := service.DeletePeminjaman("1")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "delete error")
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_GetDetailPeminjaman(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockPeminjaman := &domain.Peminjaman{
		IDPeminjaman:  "1",
		IDAnggota:     "anggota1",
		TglPinjam:     time.Now(),
		TglHrsKembali: time.Now().Add(time.Hour * 24 * 7),
		Jaminan:       "KTP",
		Details: []domain.DetailPinjam{
			{
				IDDetailpinjam: "detail1",
				IDBuku:         "buku1",
				Kondisi:        "baik",
			},
		},
	}

	mockRepo.On("GetDetailPeminjaman", "1").Return(mockPeminjaman, nil)

	peminjaman, err := service.GetDetailPeminjaman("1")

	assert.Nil(t, err)
	assert.Equal(t, mockPeminjaman, peminjaman)
	mockRepo.AssertExpectations(t)
}

func TestPeminjamanService_GetDetailPeminjaman_Error(t *testing.T) {
	mockRepo := new(mocks.IPeminjamanRepository)
	service := peminjaman.NewPeminjamanService(mockRepo)

	mockRepo.On("GetDetailPeminjaman", "1").Return(nil, errors.New("not found"))

	peminjaman, err := service.GetDetailPeminjaman("1")

	assert.NotNil(t, err)
	assert.Nil(t, peminjaman)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertExpectations(t)
}
