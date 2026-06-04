package buku

import (
	"errors"
	"regexp"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

func (s *BukuService) CreatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error) {
	var jb domain.Penerbit_Buku

	// validasi inputan nama penerbit
	if len(data.PenerbitBuku) <= 5 || len(data.PenerbitBuku) >= 255 {
		return domain.Penerbit_Buku{}, errors.New("inputan nama penerbit buku minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan alamat penerbit
	if len(data.AlamatPenerbit) <= 5 || len(data.AlamatPenerbit) >= 255 {
		return domain.Penerbit_Buku{}, errors.New("inputan alamat penerbit minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan telp penerbit
	if len(data.TelpPenerbit) <= 5 {
		return domain.Penerbit_Buku{}, errors.New("inputan telp penerbit minimal 5 karakter nomor telp")
	}

	// validasi inputan email penerbit
	match, _ := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", data.EmailPenerbit)

	if !match {
		return domain.Penerbit_Buku{}, errors.New("bukan format email")
	}

	// validasi inputan Deskripsi penerbit
	if len(data.Deskripsi) <= 5 || len(data.Deskripsi) >= 255 {
		return domain.Penerbit_Buku{}, errors.New("inputan Deskripsi penerbit minimal 5 karakter dan maksimal 255")
	}
	// end validation

	data, err := s.Repository.CreatePenerbitBuku(data)
	if err != nil {
		return jb, err
	}
	return data, nil
}

func (s *BukuService) GetAllPenerbitBuku() ([]domain.Penerbit_Buku, error) {
	data, err := s.Repository.GetAllPenerbitBuku()
	if err != nil {
		return []domain.Penerbit_Buku{}, nil
	}
	if len(data) == 0 {
		return []domain.Penerbit_Buku{}, errors.New("data kosong")
	}
	return data, nil
}

func (s *BukuService) FindPenerbitBuku(c string) ([]domain.Penerbit_Buku, error) {
	data, err := s.Repository.CariPenerbitBuku(c)
	if err != nil {
		return []domain.Penerbit_Buku{}, nil
	}
	if len(data) == 0 {
		return []domain.Penerbit_Buku{}, errors.New("data kosong")
	}
	return data, nil
}

func (s *BukuService) GetPenerbitBukuById(id string) (*domain.Penerbit_Buku, error) {
	JenisBukuById, err := s.Repository.GetPenerbitBukuById(id)
	if err != nil {
		return nil, err
	}
	if JenisBukuById.IDPenerbit == "" {
		return &domain.Penerbit_Buku{}, errors.New("data kosong")
	}
	return JenisBukuById, nil
}

func (s *BukuService) UpdatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error) {
	if s.Repository.HitungDataPenerbitBuku(data.IDPenerbit) == 0 {
		return domain.Penerbit_Buku{}, errors.New("data kosong")
	}
	// validasi inputan nama penerbit
	if len(data.PenerbitBuku) <= 5 || len(data.PenerbitBuku) >= 255 {
		return domain.Penerbit_Buku{}, errors.New("inputan nama penerbit buku minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan alamat penerbit
	if len(data.AlamatPenerbit) <= 5 || len(data.AlamatPenerbit) >= 255 {
		return domain.Penerbit_Buku{}, errors.New("inputan alamat penerbit minimal 5 karakter dan maksimal 255")
	}

	// validasi inputan telp penerbit
	if len(data.TelpPenerbit) <= 5 {
		return domain.Penerbit_Buku{}, errors.New("inputan telp penerbit minimal 5 karakter nomor telp")
	}

	// validasi inputan email penerbit
	match, _ := regexp.MatchString("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$", data.EmailPenerbit)

	if !match {
		return domain.Penerbit_Buku{}, errors.New("bukan format email")
	}

	// validasi inputan Deskripsi penerbit
	if len(data.Deskripsi) <= 5 || len(data.Deskripsi) >= 255 {
		return domain.Penerbit_Buku{}, errors.New("inputan Deskripsi penerbit minimal 5 karakter dan maksimal 255")
	}
	// end validation

	data, err := s.Repository.UpdatePenerbitBuku(data)
	if err != nil {
		return domain.Penerbit_Buku{}, err
	}
	return data, nil
}

func (s *BukuService) HapusPenerbitBuku(id string) error {
	if s.Repository.HitungDataPenerbitBuku(id) == 0 {
		return errors.New("data kosong")
	}
	if err := s.Repository.DeletePenerbitBuku(id); err != nil {
		return err
	}
	return nil
}

