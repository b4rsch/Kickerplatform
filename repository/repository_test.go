package repository

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

//TODO test failing and possible errors paths
func TestCreateUser(t *testing.T) {
	t.Run("can persist user from username and location", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected", err)
		}

		defer db.Close()
		mock.ExpectBegin()
		name := "lars"
		locationId := 1
		mock.ExpectExec("INSERT INTO users (username, locationId) VALUES ($1, $1)").
			WithArgs(name, locationId).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()
		repo := &Repository{db: db}
		_ = repo.CreateUser(name, locationId)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
