package handler

import (
	"github.com/b4rsch/Kickerplatform/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

type repositoryMock struct {}

func (repositoryMock) CreateNewMatch(_ string, _ int) (int, error) {
	return 1, nil
}

func TestCanHandleNewMatch(t *testing.T) {
	rm := repositoryMock{}
	team1 := Team{Attacker: "Lars", Defender: "Tom"}
	team2 := Team{Attacker: "Andreas", Defender: "Martin"}
	matchDetails := MatchDetails{
		Team1: team1,
		Team2: team2,
		LocationID: "2",
	}
	got, err := NewMatch(rm, matchDetails)
	if err != nil {
		t.Errorf("Got error, %s", err)
	}
	assert.Equal(t, 1, got)
}

func NewMatch(repo repository.Repo, matchDetails MatchDetails) (int, error) {}
