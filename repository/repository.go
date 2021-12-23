package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host 		= "kickerplatform-db-1"
	port 		= 5432
	user 		= "admin"
	password 	= "password123"
	dbname 		= "kickerplatformdb"
)

type Repo interface {
	CreateUser(username string, locationId int) error
}

type Repository struct {
	db *sql.DB
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


