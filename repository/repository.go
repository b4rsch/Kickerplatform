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
}

type MatchDetails struct {
	Team1      Team `json:"team1"`
	Team2      Team `json:"team2"`
	LocationID int  `json:"locationId"`
	Date       string
	ModeId     int `json:"modeId"`
}

type Team struct {
	Attacker string `json:"attacker"`
	Defender string `json:"defender"`
}

type Repository struct {
	db *sql.DB
}

func (r *Repository) CreateNewMatch(matchDetails MatchDetails) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Commit()
	stmt := `INSERT INTO matches(date, locationId, mode, team_1_attacker, team_1_defender, team_2_attacker, team_2_defender) VALUES($1, $2, $3, $4, $5, $6, $7`
	_, err = tx.Exec(
		stmt,
		matchDetails.Date,
		matchDetails.LocationID,
		matchDetails.ModeId,
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
