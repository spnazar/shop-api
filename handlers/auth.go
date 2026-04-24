package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my-secret-key")

type User struct {
	Username string
	Password string
}

var users = []User{
	{Username: "alibek", Password: "12345"},
	{Username: "admin", Password: "admin123"},
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		fmt.Fprintf(w, `{"error":"метод не разрешён"}`)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// Проверяем пользователя
	var found bool
	for _, u := range users {
		if u.Username == username && u.Password == password {
			found = true
			break
		}
	}

	if !found {
		w.WriteHeader(401)
		fmt.Fprintf(w, `{"error":"неверный логин или пароль"}`)
		return
	}

	// Создаём токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprintf(w, `{"error":"ошибка создания токена"}`)
		return
	}

	fmt.Fprintf(w, `{"token":"%s"}`, tokenString)
}

func ValidateToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	return err == nil && token.Valid
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Берём токен из заголовка
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(401)
			fmt.Fprintf(w, `{"error":"токен не предоставлен"}`)
			return
		}

		// Проверяем токен
		if !ValidateToken(tokenString) {
			w.WriteHeader(401)
			fmt.Fprintf(w, `{"error":"неверный токен"}`)
			return
		}

		// Токен валидный — пропускаем
		next(w, r)
	}
}