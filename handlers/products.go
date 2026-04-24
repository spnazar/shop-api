package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

type ProductHandler struct {
	DB *sql.DB
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := h.DB.Query("SELECT id, name, price FROM products")
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка базы данных"}`)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, price int
		var name string
		rows.Scan(&id, &name, &price)
		fmt.Fprintf(w, `{"id":%d,"name":"%s","price":%d}`+"\n", id, name, price)
	}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"метод не разрешён"}`)
		return
	}

	name := r.FormValue("name")
	price := r.FormValue("price")
	categoryID := r.FormValue("category_id")

	if name == "" || price == "" || categoryID == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error":"заполни все поля"}`)
		return
	}

	_, err := h.DB.Exec(
		"INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3)",
		name, price, categoryID,
	)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка добавления"}`)
		return
	}

	w.WriteHeader(201)
	fmt.Fprintf(w, `{"message":"товар добавлен!"}`)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"метод не разрешён"}`)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error":"укажи id товара"}`)
		return
	}

	result, err := h.DB.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка удаления"}`)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		w.WriteHeader(404)
		fmt.Fprintf(w, `{"error":"товар не найден"}`)
		return
	}

	fmt.Fprintf(w, `{"message":"товар удалён!"}`)
}