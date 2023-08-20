package config

import (
	"shop/config/migrations"
	"database/sql"
	"fmt"

	_"github.com/lib/pq"
)

const (
	Host	 = "localhost"
	Port	 = 5432
	User	 = "postgres"
	Password = "root"
	DbName	 = "Shop"
)


var (
	DB  *sql.DB
	Err error
)


func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf(
					"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
					Host, Port, User, Password, DbName,
				)

	DB, Err = sql.Open("postgres", psqlInfo)
	if Err != nil {
		panic(Err)
	}

	migrations.DbMigrate(DB)

	return DB
}