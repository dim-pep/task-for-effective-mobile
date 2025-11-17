package main

import (
	"log"
	"net/http"

	docs "github.com/dim-pep/task-for-effective-mobile/docs"
	"github.com/dim-pep/task-for-effective-mobile/internal/db"
	"github.com/dim-pep/task-for-effective-mobile/internal/web"
	"github.com/joho/godotenv"
)

// @title Subscriptions API
// @version 1.0
// @description REST-сервис для агрегации онлайн-подписок
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Ошибка чтения env файлa: %v", err)
	}

	docs.SwaggerInfo.Title = "Subscriptions API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	conn_db := db.Conn()
	db.Migrations(conn_db)

	r := web.CreateRouter()
	http.ListenAndServe(":8080", r)
}
