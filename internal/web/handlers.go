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

// 8080/subscriptions -body
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

// 8080/subscriptions{id}
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
	json.NewEncoder(w).Encode(resp)
}

// 8080/subscriptions{id}
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

// 8080/subscriptions{id}
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

// 8080/subscriptions/list
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

// 8080/subscriptions/filtred -param
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
	//json.NewEncoder(w).Encode(resp)
}
