package seed

import (
	"log"
	"time"

	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	utility "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/oklog/ulid/v2"
)

var anggota = []domain.Anggota{
	{
		IDAnggota: ulid.Make().String(),
		Username:  "budi.santoso",
		Password:  (&utility.Hash{}).HashPassword("anggota123"),
		Nama:      "Budi Santoso",
		Sex:       "Laki-laki",
		Telp:      "081234567890",
		Alamat:    "Jl. Merdeka No. 10, Jakarta Pusat",
		Email:     "budi.santoso@email.com",
		Deskripsi: "Anggota perpustakaan aktif",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		IDAnggota: ulid.Make().String(),
		Username:  "siti.rahayu",
		Password:  (&utility.Hash{}).HashPassword("anggota123"),
		Nama:      "Siti Rahayu",
		Sex:       "Perempuan",
		Telp:      "082345678901",
		Alamat:    "Jl. Sudirman No. 25, Bandung",
		Email:     "siti.rahayu@email.com",
		Deskripsi: "Anggota perpustakaan aktif",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		IDAnggota: ulid.Make().String(),
		Username:  "andi.wijaya",
		Password:  (&utility.Hash{}).HashPassword("anggota123"),
		Nama:      "Andi Wijaya",
		Sex:       "Laki-laki",
		Telp:      "083456789012",
		Alamat:    "Jl. Diponegoro No. 5, Surabaya",
		Email:     "andi.wijaya@email.com",
		Deskripsi: "Anggota perpustakaan aktif",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		IDAnggota: ulid.Make().String(),
		Username:  "dewi.kartika",
		Password:  (&utility.Hash{}).HashPassword("anggota123"),
		Nama:      "Dewi Kartika",
		Sex:       "Perempuan",
		Telp:      "084567890123",
		Alamat:    "Jl. Gajah Mada No. 8, Yogyakarta",
		Email:     "dewi.kartika@email.com",
		Deskripsi: "Anggota perpustakaan aktif",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		IDAnggota: ulid.Make().String(),
		Username:  "reza.pratama",
		Password:  (&utility.Hash{}).HashPassword("anggota123"),
		Nama:      "Reza Pratama",
		Sex:       "Laki-laki",
		Telp:      "085678901234",
		Alamat:    "Jl. Ahmad Yani No. 3, Medan",
		Email:     "reza.pratama@email.com",
		Deskripsi: "Anggota perpustakaan aktif",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

func (s *Seeds) Seed_Anggota() {
	for i := range anggota {
		err := s.Db.Model(&domain.Anggota{}).Create(&anggota[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data anggota: %v", err)
		}
	}
}
