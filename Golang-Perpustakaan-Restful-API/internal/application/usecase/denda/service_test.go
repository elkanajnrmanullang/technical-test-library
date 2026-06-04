package denda_test

import (
	"errors"
	"testing"
	"time"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/denda"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/application/usecase/denda/mocks"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDendaService_GetAllDenda(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockDenda := []domain.Denda{
		{
			IDDenda:       "1",
			JumlahDenda:   10000,
			Tglpinjam:     time.Now(),
			Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
			Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
			IDPeminjaman:  "peminjaman1",
			IDAnggota:     "anggota1",
		},
	}

	mockRepo.On("GetAllDenda").Return(mockDenda, nil)

	dendas, err := service.GetAllDenda()

	assert.Nil(t, err)
	assert.Equal(t, mockDenda, dendas)
	mockRepo.AssertExpectations(t)
}

func TestDendaService_GetAllDenda_Error(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockRepo.On("GetAllDenda").Return(nil, errors.New("database error"))

	dendas, err := service.GetAllDenda()

	assert.NotNil(t, err)
	assert.Nil(t, dendas)
	assert.EqualError(t, err, "database error")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_GetDendaByID(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockDenda := &domain.Denda{
		IDDenda:       "1",
		JumlahDenda:   10000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("GetDendaByID", "1").Return(mockDenda, nil)

	denda, err := service.GetDendaByID("1")

	assert.Nil(t, err)
	assert.Equal(t, mockDenda, denda)
	mockRepo.AssertExpectations(t)
}

func TestDendaService_GetDendaByID_Error(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockRepo.On("GetDendaByID", "1").Return(nil, errors.New("not found"))

	denda, err := service.GetDendaByID("1")

	assert.NotNil(t, err)
	assert.Nil(t, denda)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_CreateDenda(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	newDenda := &domain.Denda{
		JumlahDenda:   10000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("CreateDenda", mock.AnythingOfType("*domain.Denda")).Return(nil)

	err := service.CreateDenda(newDenda)

	assert.Nil(t, err)
	assert.NotEmpty(t, newDenda.IDDenda)
	mockRepo.AssertExpectations(t)
}

func TestDendaService_CreateDenda_Error(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	newDenda := &domain.Denda{
		JumlahDenda:   10000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("CreateDenda", mock.AnythingOfType("*domain.Denda")).Return(errors.New("create error"))

	err := service.CreateDenda(newDenda)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "create error")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_UpdateDenda(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	updatedDenda := &domain.Denda{
		IDDenda:       "1",
		JumlahDenda:   20000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("CountDendaByID", "1").Return(int64(1), nil)
	mockRepo.On("UpdateDenda", updatedDenda).Return(nil)

	err := service.UpdateDenda(updatedDenda)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDendaService_UpdateDenda_NotFound(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	updatedDenda := &domain.Denda{
		IDDenda:       "1",
		JumlahDenda:   20000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("CountDendaByID", "1").Return(int64(0), nil)

	err := service.UpdateDenda(updatedDenda)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "denda not found")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_UpdateDenda_CountError(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	updatedDenda := &domain.Denda{
		IDDenda:       "1",
		JumlahDenda:   20000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("CountDendaByID", "1").Return(int64(0), errors.New("count error"))

	err := service.UpdateDenda(updatedDenda)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "count error")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_UpdateDenda_UpdateError(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	updatedDenda := &domain.Denda{
		IDDenda:       "1",
		JumlahDenda:   20000,
		Tglpinjam:     time.Now(),
		Tglhrskembali: time.Now().Add(time.Hour * 24 * 7),
		Tglkembali:    time.Now().Add(time.Hour * 24 * 8),
		IDPeminjaman:  "peminjaman1",
		IDAnggota:     "anggota1",
	}

	mockRepo.On("CountDendaByID", "1").Return(int64(1), nil)
	mockRepo.On("UpdateDenda", updatedDenda).Return(errors.New("update error"))

	err := service.UpdateDenda(updatedDenda)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "update error")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_DeleteDenda(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockRepo.On("CountDendaByID", "1").Return(int64(1), nil)
	mockRepo.On("DeleteDenda", "1").Return(nil)

	err := service.DeleteDenda("1")

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDendaService_DeleteDenda_NotFound(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockRepo.On("CountDendaByID", "1").Return(int64(0), nil)

	err := service.DeleteDenda("1")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "denda not found")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_DeleteDenda_CountError(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockRepo.On("CountDendaByID", "1").Return(int64(0), errors.New("count error"))

	err := service.DeleteDenda("1")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "count error")
	mockRepo.AssertExpectations(t)
}

func TestDendaService_DeleteDenda_DeleteError(t *testing.T) {
	mockRepo := new(mocks.IDendaRepository)
	service := denda.NewDendaService(mockRepo)

	mockRepo.On("CountDendaByID", "1").Return(int64(1), nil)
	mockRepo.On("DeleteDenda", "1").Return(errors.New("delete error"))

	err := service.DeleteDenda("1")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "delete error")
	mockRepo.AssertExpectations(t)
}
