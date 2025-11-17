package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dim-pep/task-for-effective-mobile/internal/config"
	"github.com/dim-pep/task-for-effective-mobile/internal/db"
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

// postSub godoc
// @Summary Создать подписку
// @Description Создаёт запись о подписке.
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param input body config.Subscriptions true "Данные подписки"
// @Success 201 {string} string "created"
// @Failure 400 {string} string "Неправильный JSON/валидация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions [post]
func postSub(w http.ResponseWriter, r *http.Request) { //Create
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	dec := json.NewDecoder(r.Body)

	var req config.Subscriptions
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Неправильный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.ServiceName == "" {
		http.Error(w, "service_name не указан", http.StatusBadRequest)
		return
	}
	if req.Price <= 0 {
		http.Error(w, "price равен 0 или меньше нуля", http.StatusBadRequest)
		return
	}
	if req.UserID == "" {
		http.Error(w, "user_id не указан", http.StatusBadRequest)
		return
	}
	_, err := time.Parse("2006-01", req.StartDate)
	if err != nil {
		http.Error(w, "указан не верный start_date", http.StatusBadRequest)
		return
	}

	err = db.CreateSub(req)
	if err != nil {
		http.Error(w, "Ошибка при создание объекта базы данных", http.StatusInternalServerError)
		return
	}
	log.Println("Объект успешно создан")
	w.WriteHeader(http.StatusCreated)
}

// getSub godoc
// @Summary Получить подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "ID подписки"
// @Success 200 {object} config.Subscriptions
// @Failure 400 {string} string "Объекта с таким id не существует"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [get]
func getSub(w http.ResponseWriter, r *http.Request) { //Read
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := chi.URLParam(r, "id")
	resp, err := db.GetSub(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Объекта с таким id не существует", http.StatusBadRequest)
			return
		}
		http.Error(w, "Ошибка при чтении объекта базы данных", http.StatusInternalServerError)
		return
	}
	log.Println("Объект успешно получен")
	w.WriteHeader(http.StatusOK)
	resp.StartDate = resp.StartDate[0:7]
	json.NewEncoder(w).Encode(resp)
}

// putSub godoc
// @Summary Обновить подписку по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "ID подписки"
// @Param input body config.Subscriptions true "Новые данные подписки"
// @Success 200 {string} string "updated"
// @Failure 204 {string} string "Объекта с таким id не существует"
// @Failure 400 {string} string "Неправильный JSON/валидация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [put]
func putSub(w http.ResponseWriter, r *http.Request) { //Update
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	dec := json.NewDecoder(r.Body)

	var req config.Subscriptions
	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Неправильный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.ServiceName == "" {
		http.Error(w, "service_name не указан", http.StatusBadRequest)
		return
	}
	if req.Price <= 0 {
		http.Error(w, "price равен 0 или меньше нуля", http.StatusBadRequest)
		return
	}
	if req.UserID == "" {
		http.Error(w, "user_id не указан", http.StatusBadRequest)
		return
	}
	_, err := time.Parse("2006-01", req.StartDate)
	if err != nil {
		http.Error(w, "указан не верный start_date", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	num, err := db.UpdateSub(id, req)
	if err != nil {
		http.Error(w, "Ошибка при изменении объекта базы данных", http.StatusInternalServerError)
		return
	}
	if num == 0 {
		http.Error(w, "Объекта с таким id не существует", http.StatusNoContent)
		return
	}
	log.Println("Объект успешно обновлён")
	w.WriteHeader(http.StatusOK)
}

// delSub godoc
// @Summary Удалить подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path string true "ID подписки"
// @Success 200 {string} string "deleted"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [delete]
func delSub(w http.ResponseWriter, r *http.Request) { //Delete
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id := chi.URLParam(r, "id")
	err := db.DelSub(id)
	if err != nil {
		http.Error(w, "Ошибка при удаления объекта базы данных", http.StatusInternalServerError)
		return
	}
	log.Println("Объект успешно удалён")
	w.WriteHeader(http.StatusOK)
}

// getListSub godoc
// @Summary Список подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {array} config.Subscriptions
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/list [get]
func getListSub(w http.ResponseWriter, r *http.Request) { //List?????
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp, err := db.GetSubsId()
	if err != nil {
		http.Error(w, "Ошибка при удаления объекта базы данных", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Лист объектов успешно получен")
	json.NewEncoder(w).Encode(resp)
}

// getSumFiltredListOfSub godoc
// @Summary Суммарная стоимость подписок за период
// @Description Возвращает суммарную стоимость за период с фильтрацией по user_id и service_name. Даты в формате YYYY-MM.
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param input body config.FilterRequest true "Фильтр: service_name, user_id, start_date (YYYY-MM), end_date (YYYY-MM)"
// @Success 200 {object} map[string]int "Пример: {\"total_cost\":1200}"
// @Failure 400 {string} string "Неправильный JSON/валидация"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/sum [post]
func getSumFiltredListOfSub(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	dec := json.NewDecoder(r.Body)
	var req config.FilterRequest

	if err := dec.Decode(&req); err != nil {
		http.Error(w, "Неправильный JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.ServiceName == "" {
		http.Error(w, "service_name не указан", http.StatusBadRequest)
		return
	}
	if req.UserID == "" {
		http.Error(w, "user_id не указан", http.StatusBadRequest)
		return
	}
	_, err := time.Parse("2006-01", req.StartDate)
	if err != nil {
		http.Error(w, "указан не верный start_date", http.StatusBadRequest)
		return
	}
	_, err = time.Parse("2006-01", req.EndDate)
	if err != nil {
		http.Error(w, "указан не верный end_date", http.StatusBadRequest)
		return
	}

	sum, err := db.GetSumFiltredSubs(req.StartDate, req.EndDate, req.UserID, req.ServiceName)
	if err != nil {
		http.Error(w, "Ошибка выполнении запроса", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{
		"total_cost": sum,
	})
}
