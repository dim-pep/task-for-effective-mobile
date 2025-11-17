package main

import (
	"log"
	"net/http"

	"github.com/dim-pep/task-for-effective-mobile/internal/db"
	"github.com/dim-pep/task-for-effective-mobile/internal/web"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Ошибка чтения env файлa: %v", err)
	}

	conn_db := db.Conn()
	db.Migrations(conn_db) // psql -U postgres -d subscriptions_db

	r := web.CreateRouter()
	http.ListenAndServe(":8080", r)
}
