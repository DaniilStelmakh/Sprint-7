package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestMainHandlerWhenCountMoreThanTotal для проверки, если count передано больше ожидаемого
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// количество значений кафе в городе
	totalCount := 4
	// GET запрос к севреру
	req := httptest.NewRequest("GET", "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)

	assert.Len(t, strings.Split(responseRecorder.Body.String(), ","), totalCount)

}

// TestMainHandlerWhenResponseNotSupported для проверки, если city не поддерживается
func TestMainHandlerWhenResponseNotSupported(t *testing.T) {
	// GET запрос к серверу
	req := httptest.NewRequest("GET", "/cafe?count=4&city=omsk", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	assert.NotEqual(t, cafeList, req.URL.Query().Get("city"))

	assert.Equal(t, "wrong city value", responseRecorder.Body.String())
}

// TestMainHandlerWhenResponseBodyNotEmpty для проверки, что тело ответа не пустое
func TestMainHandlerWhenResponseBodyNotEmpty(t *testing.T) {
	// GET запрос к серверу
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String())
}
