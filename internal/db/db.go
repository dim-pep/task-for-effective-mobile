package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	pg_user := os.Getenv("POSTGRES_USER")
	pg_pass := os.Getenv("POSTGRES_PASSWORD")
	pg_port := os.Getenv("APP_PORT")

	dsn := "postgres://" + pg_user + ":" + pg_pass + "@localhost:" + pg_port + "/subscriptions_db?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	log.Println("Успешно подключился к базе данных")
	return db
}

func test() {
	log.Println("Privet")
}
