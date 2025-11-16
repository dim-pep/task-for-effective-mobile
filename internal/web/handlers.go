package web

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// 8080/subscriptions
func postSub(w http.ResponseWriter, r *http.Request) { //Create
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Получить запись"))
}

// 8080/subscriptions{id}
func getSub(w http.ResponseWriter, r *http.Request) { //Read
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id должен быть числом", http.StatusBadRequest)
		return
	}

	fmt.Println(id)

	w.Write([]byte("Получить запись"))
}

// 8080/subscriptions{id}
func putSub(w http.ResponseWriter, r *http.Request) { //Update
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Получить запись"))
}

// 8080/subscriptions{id}
func delSub(w http.ResponseWriter, r *http.Client) { //Delete
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Получить запись"))
}

// 8080/subscriptions
func listSub(w http.ResponseWriter, r *http.Client) { //List?????
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Получить запись"))
}

// 8080/subscriptions -param
func filtredListSub(w http.ResponseWriter, r *http.Client) { //подсчета суммарной стоимости всех подписок за выбранный период с фильтрацией по id пользователя и названию подписки
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Получить запись"))
}

func test() {
	fmt.Println("Privet")
}
