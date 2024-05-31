package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	uri := os.Getenv("DATABASE_URI")
	// connStr := "user='rmr-adm' password=UarLS1xD3cnR host=ep-silent-dew-a254ql5e.eu-central-1.pg.koyeb.app dbname='campamento'"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
