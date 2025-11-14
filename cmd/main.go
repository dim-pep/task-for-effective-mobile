package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://postgres:mysecretpassword@localhost:5432/subscriptions_db?sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	fmt.Println("подключился к базе данных")

	sqlFile, err := os.ReadFile("../migrations/001_init.sql")
	if err != nil {
		log.Fatalf("Ошибка чтения SQL-файла: %v", err)
	}

	sqlScript := string(sqlFile)
	result := db.Exec(sqlScript)

	if result.Error != nil {
		log.Fatalf("Ошибка выполнения SQL-скрипта: %v", result.Error)
	}

	fmt.Println("Таблица успешно создана")

}
