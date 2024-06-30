package main

import (
	"log"
	"os"

	"github.com/ChiragRajput101/rest-api/config"
	"github.com/ChiragRajput101/rest-api/db"
	mysqlCfg "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg := mysqlCfg.Config {
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := db.NewMySQLStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	// m.Up() -> create migration
	// m.Down() -> delete

	command := os.Args[(len(os.Args)) - 1]

	if command == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		}
	}

	if command == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange{
			log.Fatal(err)
		}
	}
}