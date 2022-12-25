package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDefaultPage(t *testing.T) {
	t.Parallel()
	w := httptest.NewRecorder()

	ctx, r := gin.CreateTestContext(w)

	setupRouter(r)

	req, err := http.NewRequestWithContext(ctx, "GET", "/", nil)
	if err != nil {
		t.Errorf("got error: %s", err)
	}
	//запускаем запрос через наш роутер, записывая результат на w
	r.ServeHTTP(w, req)
	want := strconv.Itoa(http.StatusOK)
	got := strconv.Itoa(w.Code)
	if want != got {
		t.Fatalf("expected response code %v, got %v", want, got)
	}

	body := w.Body.String()

	want = "Hello, Nick!"
	got = strings.Trim(body, " \r\n")
	if want != got {
		t.Fatalf("expected response body '%s', got '%s'", want, body)
	}
}
