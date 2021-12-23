package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "kickerplatform-db-1"
	port     = 5432
	user     = "admin"
	password = "password123"
	dbname   = "kickerplatformdb"
)

type Repo interface {
	CreateUser(username string, locationId int) error
	CreateNewMatch(matchDetails MatchDetails) (int, error)
	WriteGameOfMatch(matchId string, payload GameDetails) error
	GetPlayedMatchesFor(username string)(int, error)
}

type MatchDetails struct {
	Team1      Team `json:"team1"`
	Team2      Team `json:"team2"`
	LocationID int  `json:"locationId"`
	Date       string
	BestOf     int `json:"best_of"`
}

type Team struct {
	Attacker string `json:"attacker"`
	Defender string `json:"defender"`
}

type GameDetails struct {
	Team1 Team `json:"team1"`
	Team2 Team `json:"team2"`
	ScoreTeam1 int `json:"score_team_1"`
	ScoreTeam2 int `json:"score_team_2"`
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetPlayedMatchesFor(username string) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Commit()

	rows, _ := tx.Query("SELECT COUNT(*) FROM matches WHERE team_1_player_1 = $1 or team_1_player_2 =  $1")
	var result int
	for rows.Next() {
		_ = rows.Scan(&result)
	}

	return result, nil
}

func (r *Repository) CreateNewMatch(matchDetails MatchDetails) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Commit()
	//TODO team players as ID with reference to users
	stmt := `INSERT INTO matches(date, locationId, best_of, team_1_player_1, team_1_player_2, team_2_player_1, team_2_player_2) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(
		stmt,
		matchDetails.Date,
		matchDetails.LocationID,
		matchDetails.BestOf,
		matchDetails.Team1.Attacker,
		matchDetails.Team1.Defender,
		matchDetails.Team2.Attacker,
		matchDetails.Team2.Defender,
	)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func (r *Repository) CreateUser(name string, locationId int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	stmt := `INSERT INTO users(username, locationId) VALUES($1, $2)`
	_, err = tx.Exec(stmt, name, locationId)
	if err != nil {
		return err
	}
	return nil
}

// NewRepository initializes repository with db
func NewRepository() *Repository {
	r := Repository{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("could not connect to database")
		panic(err)
	}
	r.db = db
	return &r
}

func (r *Repository) WriteGameOfMatch(matchId string, game GameDetails) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	stmt := `INSERT INTO game(match_id, points_team_1, points_team_2, team_1_attacker, team_1_defender, team_2_attacker, team_2_defender) VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err = tx.Exec(
		stmt,
		matchId,
		game.ScoreTeam1,
		game.ScoreTeam2,
		game.Team1.Attacker,
		game.Team1.Defender,
		game.Team2.Attacker,
		game.Team2.Defender,
		)
	if err != nil {
		return err
	}
	return nil
}
