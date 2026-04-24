package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	godotenv.Load()

	connStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return nil
	}

	err = database.Ping()
	if err != nil {
		fmt.Println("Не могу подключиться:", err)
		return nil
	}

	fmt.Println("Подключились к PostgreSQL!")
	return database
}

func Init(database *sql.DB) {
	database.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	)`)

	database.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		price INT NOT NULL,
		category_id INT REFERENCES categories(id)
	)`)

	var count int
	database.QueryRow("SELECT COUNT(*) FROM categories").Scan(&count)
	if count == 0 {
		database.Exec("INSERT INTO categories (name) VALUES ($1)", "Телефоны")
		database.Exec("INSERT INTO categories (name) VALUES ($1)", "Ноутбуки")
		database.Exec("INSERT INTO categories (name) VALUES ($1)", "Наушники")

		database.Exec("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3)", "iPhone", 500000, 1)
		database.Exec("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3)", "Samsung", 400000, 1)
		database.Exec("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3)", "MacBook", 1200000, 2)
		database.Exec("INSERT INTO products (name, price, category_id) VALUES ($1, $2, $3)", "AirPods", 150000, 3)
		fmt.Println("Данные добавлены!")
	}
}