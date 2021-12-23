package handler

type MatchDetails struct {
	Team1 Team `json:"team1"`
	Team2 Team `json:"team2"`
	LocationID string `json:"locationId"`
}
type Team struct {
	Attacker string `json:"attacker"`
	Defender string `json:"defender"`
}

