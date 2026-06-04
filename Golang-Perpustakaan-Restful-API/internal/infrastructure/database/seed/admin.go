package seed

import (
	"log"
	"time"

	domain "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/core/domain"
	utility "github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/utility"
	"github.com/oklog/ulid/v2"
)

var hash utility.Hash
var admin = []domain.Pegawai{
	{
		IDPegawai: ulid.Make().String(),
		Username:  "admin",
		Password:  (&utility.Hash{}).HashPassword("admin"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
	{
		IDPegawai: ulid.Make().String(),
		Username:  "afrizal",
		Password:  (&utility.Hash{}).HashPassword("admin"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	},
}

func (s *Seeds) Seed_admin() {
	// masukkan ke db
	for i, _ := range admin {
		err := s.Db.Model(&domain.Pegawai{}).Create(&admin[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data admin: %v", err)
		}
	}
}
