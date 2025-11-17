package db

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/dim-pep/task-for-effective-mobile/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Gorm *gorm.DB

func Conn() *gorm.DB {
	pg_user := os.Getenv("POSTGRES_USER")
	pg_pass := os.Getenv("POSTGRES_PASSWORD")
	pg_port := os.Getenv("DB_PORT")
	pg_host := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/subscriptions_db?sslmode=disable", pg_user, pg_pass, pg_host, pg_port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto;")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	log.Println("Успешно подключился к базе данных")
	Gorm = db
	return db
}

func CreateSub(sub config.Subscriptions) error {
	sub.StartDate += "-01"
	res := Gorm.Create(sub)
	if res.Error != nil {
		log.Panicf("Ошибка при создание объекта базы данных: " + res.Error.Error())
		return res.Error
	}
	return nil
}

func GetSub(id string) (config.Subscriptions, error) {
	var ans config.Subscriptions
	res := Gorm.First(&ans, "id = "+"'"+id+"'")
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return config.Subscriptions{}, gorm.ErrRecordNotFound
		}
		log.Panicf("Ошибка при чтении объекта базы данных: " + res.Error.Error())
		return config.Subscriptions{}, res.Error
	}
	return ans, nil
}

func UpdateSub(id string, sub config.Subscriptions) (int64, error) {
	sub.StartDate += "-01"
	res := Gorm.Model(&config.Subscriptions{}).Where("id = ?", id).Updates(sub)
	if res.Error != nil {
		log.Panicf("Ошибка при изменение объекта базы данных: " + res.Error.Error())
		return 0, res.Error
	}
	return res.RowsAffected, nil
}

func DelSub(id string) error {
	res := Gorm.Unscoped().Delete(&config.Subscriptions{}, "id = "+"'"+id+"'")
	if res.Error != nil {
		log.Panicf("Ошибка при удалении объекта базы данных: " + res.Error.Error())
		return res.Error
	}
	return nil
}

func GetSubsId() (map[string][]string, error) {
	var ans []string
	res := Gorm.Model(&config.Subscriptions{}).Pluck("id", &ans)
	if res.Error != nil {
		log.Panicf("Ошибка при чтении объекта базы данных: " + res.Error.Error())
		return map[string][]string{}, res.Error
	}
	resp := map[string][]string{
		"id": ans,
	}
	return resp, nil
}

func GetSumFiltredSubs(periodStart, periodEnd, userID, serviceName string) (int, error) {
	tx := Gorm.Model(&config.Subscriptions{})
	tx.Where("start_date <= ?", periodEnd+"-01")
	tx.Where("user_id = ?", userID)
	tx.Where("start_date >= ?", periodStart+"-01")
	Query := "SUM(price)"
	var ans int
	res := tx.Select(Query).Scan(&ans)
	if res.Error != nil {
		log.Panicf("Ошибка при выполнении запроса: " + res.Error.Error())
		return 0, res.Error
	}
	return ans, nil
}
