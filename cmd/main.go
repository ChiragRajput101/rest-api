package main

import (
	"database/sql"
	"log"
	"fmt"

	"github.com/ChiragRajput101/rest-api/cmd/api"
	"github.com/ChiragRajput101/rest-api/config"
	"github.com/ChiragRajput101/rest-api/db"
	"github.com/go-sql-driver/mysql"
)


// Ping verifies a connection to the database is still alive, 
// establishing a connection if necessary.

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db successfully conn")
} 

// instantiate the server
func main() {	

	cfg := mysql.Config {
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	}

	db, err := db.NewMySQLStorage(cfg)

	if err != nil {
		log.Fatal(err)
	}

	// Ping the DB
	initStorage(db) 

	server := api.InitServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}