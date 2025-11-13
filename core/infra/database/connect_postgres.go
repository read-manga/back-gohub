package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectPostgres() {
	connStr := "host=localhost port=5432 user=user password=password123 dbname=users sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Erro ao conectar:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Erro ao testar conexão:", err)
	}

	DB = db
	fmt.Println("Conectado ao PostgreSQL!")
}

func GetDB() *sql.DB {
	return DB
}
