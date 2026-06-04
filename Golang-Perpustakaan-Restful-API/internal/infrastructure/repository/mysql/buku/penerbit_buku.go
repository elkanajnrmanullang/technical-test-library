package buku

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

func (q *BukuRepository) GetAllPenerbitBuku() ([]domain.Penerbit_Buku, error) {
	var data, result []domain.Penerbit_Buku
	if err := q.db.Find(&data).Error; err != nil {
		return nil, err
	}
	result = append(result, data...)
	return result, nil
}

func (q *BukuRepository) CariPenerbitBuku(c string) ([]domain.Penerbit_Buku, error) {
	var data, result []domain.Penerbit_Buku
	if err := q.db.Where("penerbit_buku LIKE ? OR email_penerbit LIKE ? or deskripsi LIKE ?", "%"+c+"%", "%"+c+"%", "%"+c+"%").Find(&data).Error; err != nil {
		return nil, err
	}
	result = append(result, data...)
	return result, nil
}

func (q *BukuRepository) GetPenerbitBukuById(id string) (*domain.Penerbit_Buku, error) {
	var data domain.Penerbit_Buku
	if err := q.db.First(&data, "id_penerbit = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (q *BukuRepository) CreatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error) {
	var jb domain.Penerbit_Buku
	if err := q.db.Save(&data).Error; err != nil {
		return jb, err
	}
	return data, nil
}

func (q *BukuRepository) UpdatePenerbitBuku(data domain.Penerbit_Buku) (domain.Penerbit_Buku, error) {
	if err := q.db.Where("id_penerbit = ?", &data.IDPenerbit).Updates(&data).Error; err != nil {
		return domain.Penerbit_Buku{}, err
	}
	return data, nil
}

func (q *BukuRepository) DeletePenerbitBuku(id string) error {
	var data domain.Penerbit_Buku
	if err := q.db.Where("id_penerbit = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *BukuRepository) HitungDataPenerbitBuku(id string) int64 {
	var count int64
	if err := q.db.Model(&domain.Penerbit_Buku{}).Where("id_penerbit = ?", id).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return count
}

