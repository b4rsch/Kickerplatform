package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type repositoryMock struct {}

func (repositoryMock) CreateUser(_ string, _ int) error{
	return nil
}

func TestCanCreateUser(t *testing.T) {
	rm := &repositoryMock{}
	router := setupRouter(rm)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/user/lars/1", nil)

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCanCreateMatch(t *testing.T) {
	rm := &repositoryMock{}
	router := setupRouter(rm)
	w := httptest.NewRecorder()
	payload := `{"team1": {"attacker":"lars","defender":"tom"},"team2":{"attacker":"andreas","defender":"martin"},"locationId":"2"}`
	r, _ := http.NewRequest("POST", "/match/new", bytes.NewBuffer([]byte(payload)))

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}
