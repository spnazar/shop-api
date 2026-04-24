# Shop API

REST API для магазина написанный на Go с использованием PostgreSQL.

## Технологии

- Go
- PostgreSQL
- net/http

## Установка

1. Клонируй репозиторий
git clone https://github.com/ТВОЙ_НИКНЕЙМ/shop-api.git
cd shop-api

2. Создай файл .env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=твой_пароль
DB_NAME=shop

3. Установи зависимости
go mod tidy

4. Запусти
go run main.go

## API endpoints

| Метод | Адрес | Описание |
|-------|-------|----------|
| GET | /products | Все товары |
| GET | /categories | Все категории |
| POST | /products/create | Добавить товар |
| DELETE | /products/delete?id=1 | Удалить товар |

## Примеры запросов

Получить все товары:
GET http://localhost:8080/products

Добавить товар:
POST http://localhost:8080/products/create
name=iPhone&price=500000&category_id=1

Удалить товар:
DELETE http://localhost:8080/products/delete?id=1

## Авторизация

Получить токен:
POST /login
username=alibek&password=12345

Использовать токен:
GET /products
Authorization: ВАШ_ТОКЕН

## Структура проекта

shop-api/
├── main.go
├── db/
│   └── db.go
├── handlers/
│   ├── products.go
│   └── categories.go
├── models/
│   └── models.go
└── .env
