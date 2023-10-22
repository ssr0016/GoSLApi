package config

import (
	"database/sql"
	"fmt"
	"CrudApi/helper"
	"github.com/lib/pq" // Postgres golang driver
	"github.com/rs/zerolog/log"
)

const(
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "postgres"
	dbname = "test"
)

func DatabaseConnection() *sql.DB{
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", sqlInfo)
	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connection to database!!")

	return db
}