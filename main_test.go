package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Тест на успешный GET запрос
func TestGetUserHandler_Success(t *testing.T) {
	// Сбрасываем счетчик перед тестом, чтобы ожидать ID = 1
	globalUserID = 1

	// Создаем фейковый HTTP GET запрос к урлу /user
	req, err := http.NewRequest(http.MethodGet, "/user", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	// ResponseRecorder заменяет реальный http.ResponseWriter и записывает ответ в память
	rr := httptest.NewRecorder()

	// Напрямую вызываем наш обработчик
	GetUserHandler(rr, req)

	// 1. Проверяем HTTP статус-код
	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался статус 200, получили %d", rr.Code)
	}

	// 2. Проверяем Content-Type заголовок
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Ожидался Content-Type 'application/json', получили '%s'", contentType)
	}

	// 3. Декодируем тело ответа и проверяем ID
	var user User
	if err := json.NewDecoder(rr.Body).Decode(&user); err != nil {
		t.Fatalf("Не удалось распарсить JSON ответа: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("Ожидался ID = 1, получили %d", user.ID)
	}
}

// Тест на отправку неверного метода (например, POST)
func TestGetUserHandler_MethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/user", nil)
	if err != nil {
		t.Fatalf("Не удалось создать запрос: %v", err)
	}

	rr := httptest.NewRecorder()
	GetUserHandler(rr, req)

	// Проверяем, что метод POST заблокирован (код 405)
	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Ожидался статус 405 Method Not Allowed, получили %d", rr.Code)
	}
}
