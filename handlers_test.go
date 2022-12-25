package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHomePage(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(w)
	setupRouter(r, freshDb(t))
	req, err := http.NewRequestWithContext(ctx, "GET", "/products/", nil)
	if err != nil {
		t.Errorf("Поймали ошибку: %s", err)
	}
	r.ServeHTTP(w, req)
	if http.StatusOK != w.Code {
		t.Fatalf("Ожидали код ответа %d, получили %d", http.StatusOK, w.Code)
	}
	got := w.Body.String()
	want := "<h2>Мои изделия</h2>"

	if !strings.Contains(got, want) {
		t.Fatalf("Ожидали содержание ответа '%s', получили '%s'", want, got)
	}
}
