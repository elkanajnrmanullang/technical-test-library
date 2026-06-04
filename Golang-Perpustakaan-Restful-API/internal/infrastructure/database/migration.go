package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/internal/infrastructure/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// proses migrasi
func Migrate() {
	configDB := map[string]string{
		"DB_Username": config.Config("DB_USERNAME"),
		"DB_Password": config.Config("DB_PASSWORD"),
		"DB_Port":     config.Config("DB_PORT"),
		"DB_Host":     config.Config("DB_ADDRESS"),
		"DB_Name":     config.Config("DB_NAME"),
	}
	openConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])
	db, _ := sql.Open("mysql", openConnection)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, e := migrate.NewWithDatabaseInstance(
		"file://internal/infrastructure/database/migration",
		"mysql",
		driver,
	)
	if e != nil {
		log.Fatal(e)
	}
	err := m.Up()
	if err != nil {
		log.Fatal(err)
	}
}
