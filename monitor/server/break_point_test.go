package main

import(
	"net/http/httptest"
	"net/http"
	"testing"
	"fmt"
)

func TestBp(t *testing.T){
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/market/bp", nil)
	r.ServeHTTP(w, req)

	fmt.Println(w.Body.String())
}