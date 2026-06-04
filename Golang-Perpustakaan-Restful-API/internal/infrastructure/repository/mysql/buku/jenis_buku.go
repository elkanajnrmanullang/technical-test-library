package buku

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
)

func (q *BukuRepository) GetAllJenisBuku() ([]domain.Jenis_Buku, error) {
	var data, result []domain.Jenis_Buku
	if err := q.db.Find(&data).Error; err != nil {
		return nil, err
	}
	// for _, v := range data {
	result = append(result, data...)
	// }
	return result, nil
}

func (q *BukuRepository) CariJenisBuku(c string) ([]domain.Jenis_Buku, error) {
	var data, result []domain.Jenis_Buku
	if err := q.db.Where("jenis_buku LIKE ? OR deskripsi LIKE ?", "%"+c+"%", "%"+c+"%").Find(&data).Error; err != nil {
		return nil, err
	}
	result = append(result, data...)
	return result, nil
}

func (q *BukuRepository) GetJenisBukuById(id string) (*domain.Jenis_Buku, error) {
	var data domain.Jenis_Buku
	if err := q.db.First(&data, "id_jenis = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (q *BukuRepository) CreateJenisBuku(data domain.Jenis_Buku) (domain.Jenis_Buku, error) {
	var jb domain.Jenis_Buku
	if err := q.db.Save(&data).Error; err != nil {
		return jb, err
	}
	return data, nil
}

func (q *BukuRepository) UpdateJenisBuku(data domain.Jenis_Buku) (domain.Jenis_Buku, error) {
	if err := q.db.Where("id_jenis = ?", &data.IDJenis).Updates(&data).Error; err != nil {
		return domain.Jenis_Buku{}, err
	}
	return data, nil
}

func (q *BukuRepository) DeleteJenisBuku(id string) error {
	var data domain.Jenis_Buku
	if err := q.db.Where("id_jenis = ?", id).Delete(&data).Error; err != nil {
		return err
	}
	return nil
}

func (q *BukuRepository) HitungDataJenisBuku(id string) int64 {
	var count int64
	if err := q.db.Model(&domain.Jenis_Buku{}).Where("id_jenis = ?", id).Count(&count).Error; err != nil {
		log.Println(err)
	}
	return count
}

