package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

type CategoryHandler struct {
	DB *sql.DB
}

func (h *CategoryHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := h.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка базы данных"}`)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		fmt.Fprintf(w, `{"id":%d,"name":"%s"}`+"\n", id, name)
	}
}