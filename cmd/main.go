package main

import (
	"database/sql"
	"fmt"
	"github.com/AbnerCMoura/REST_API_MYSQL/cmd/api"
	"github.com/AbnerCMoura/REST_API_MYSQL/config"
	"github.com/AbnerCMoura/REST_API_MYSQL/db"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {
	dbase, err := db.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DbUser,
		Passwd: config.Envs.DbPassword,
		Addr:   config.Envs.DbAddress,
		DBName: config.Envs.DbName,
	})
	if err != nil {
		log.Fatal(err)
	}
	dbase.SetConnMaxLifetime(time.Minute * 3)
	dbase.SetMaxOpenConns(10)
	dbase.SetMaxIdleConns(10)

	initStorage(dbase)

	server := api.NewApiServer(":8090", dbase)
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados!")
}
