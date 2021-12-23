package main

import (
	"bytes"
	"github.com/b4rsch/Kickerplatform/repository"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type repositoryMock struct{}

func (repositoryMock) CreateUser(_ string, _ int) error {
	return nil
}
func (repositoryMock) CreateNewMatch(matchDetails repository.MatchDetails) (int, error) {
	return 1, nil
}
func (repositoryMock) WriteGameOfMatch(matchId string, payload repository.GameDetails) error{
	return nil
}
func (repositoryMock) GetPlayedMatchesFor(username string) (int, error) {
	return 1, nil
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
	payload := `{
"team1": {"attacker":"lars","defender":"tom"},
"team2":{"attacker":"andreas","defender":"martin"},
"locationId":2,
"best_of": 1,
"date": "2021-12-13"
}`
	r, _ := http.NewRequest("POST", "/match/new", bytes.NewBuffer([]byte(payload)))

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `"{\"matchId\": 1}"`, w.Body.String())
}

func TestCanEnterResultForGame(t *testing.T) {
	rm := &repositoryMock{}
	router := setupRouter(rm)
	w := httptest.NewRecorder()
	payload := `{
"team1": {"attacker":"lars","defender":"tom"},
"team2":{"attacker":"andreas","defender":"martin"},
"score_team_1": 10,
"score_team_2": 9
}`
	r, _ := http.NewRequest("POST", "/game/1", bytes.NewBuffer([]byte(payload)))

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCanGetUserStatisticsByUserId(t *testing.T) {
	rm := &repositoryMock{}
	router := setupRouter(rm)
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/user/lars", nil)

	router.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"matches_played": 5}`, w.Body.String())
}
