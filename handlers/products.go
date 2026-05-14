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

	categoryID := r.URL.Query().Get("category_id")
	var rows *sql.Rows
	var err error

	if categoryID != "" {
		rows, err = h.DB.Query("SELECT id, name, price, category_id FROM products WHERE category_id = $1", categoryID)
	} else {
		rows, err = h.DB.Query("SELECT id, name, price, category_id FROM products")
	}

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка базы данных"}`)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id, price, catID int
		var name string
		rows.Scan(&id, &name, &price, &catID)
		fmt.Fprintf(w, `{"id":%d,"name":"%s","price":%d,"category_id":%d}`+"\n", id, name, price, catID)
	}
}

func (h *ProductHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.PathValue("id")

	var productId, price, catID int
	var name string
	err := h.DB.QueryRow("SELECT id, name, price, category_id FROM products WHERE id = $1", id).Scan(&productId, &name, &price, &catID)
	
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			fmt.Fprintf(w, `{"error":"товар не найден"}`)
		} else {
			w.WriteHeader(500)
			fmt.Fprintf(w, `{"error":"ошибка базы данных"}`)
		}
		return
	}

	fmt.Fprintf(w, `{"id":%d,"name":"%s","price":%d,"category_id":%d}`+"\n", productId, name, price, catID)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Метод теперь проверяется на уровне роутера (mux)

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

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	name := r.FormValue("name")
	price := r.FormValue("price")
	categoryID := r.FormValue("category_id")

	if name == "" || price == "" || categoryID == "" {
		w.WriteHeader(400)
		fmt.Fprintf(w, `{"error":"заполни все поля"}`)
		return
	}

	result, err := h.DB.Exec(
		"UPDATE products SET name = $1, price = $2, category_id = $3 WHERE id = $4",
		name, price, categoryID, id,
	)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка обновления"}`)
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		w.WriteHeader(404)
		fmt.Fprintf(w, `{"error":"товар не найден"}`)
		return
	}

	fmt.Fprintf(w, `{"message":"товар обновлён!"}`)
}

func (h *ProductHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// Метод и наличие ID теперь гарантируются роутером
	id := r.PathValue("id")

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