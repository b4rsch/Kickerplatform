package handler

import (
	"github.com/b4rsch/Kickerplatform/repository"
	"strconv"
)

func UserStatisticsFor(username string, repo repository.Repo) (map[string]string, error) {
	var userStatistics map[string]string
	matchesPlayed, err := repo.GetPlayedMatchesFor(username)
	if err != nil {
		return userStatistics, err
	}
	userStatistics["matchesPlayed"] = strconv.Itoa(matchesPlayed)
	return userStatistics, nil
}
