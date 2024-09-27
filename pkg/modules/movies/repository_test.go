package movies_test

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"regexp"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	aliasMovies "movies/definitions/movies" // Alias for the package
	"movies/pkg/modules/movies"
)

func TestRepository_GetByID(t *testing.T) {
	// Setup SQLMock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	// Setup Gorm DB with SQLMock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Initialize the repository with the mock database
	repo := movies.Repository{Db: gormDB}

	t.Run("Success - Movie Found", func(t *testing.T) {
		// Define the expected movie result
		expectedMovie := aliasMovies.Movie{
			Id:       "1",
			Name:     "Inception",
			Genre:    "Sci-Fi",
			Year:     "2010",
			Overview: "A thief who steals corporate secrets...",
		}

		movieId, err := strconv.ParseInt(expectedMovie.Id, 10, 64) // base 10, 64-bit integer
		if err != nil {
			log.Fatalf("Error converting string to int64: %v", err)
		}

		// Mock the SQL query
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "movies" WHERE "movies"."id" = $1 ORDER BY "movies"."id" LIMIT $2`)).
			WithArgs(movieId, 1).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name", "genre", "year", "overview"}).
					AddRow(expectedMovie.Id, expectedMovie.Name, expectedMovie.Genre, expectedMovie.Year, expectedMovie.Overview))

		// Call the repository function
		actualMovie, err := repo.GetByID(1)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedMovie, actualMovie)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error - Movie Not Found", func(t *testing.T) {
		// Mock the SQL query to return no rows
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "movies" WHERE "movies"."id" = $1 ORDER BY "movies"."id" LIMIT $2`)).
			WithArgs(2, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		// Call the repository function
		_, err := repo.GetByID(2)

		// Assertions
		assert.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("movie with ID %d not found", 2))
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Error - Database Error", func(t *testing.T) {
		// Mock the SQL query to return a database error
		dbError := errors.New("database error")
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "movies" WHERE "movies"."id" = $1 ORDER BY "movies"."id" LIMIT $`)).
			WithArgs(3, 1).
			WillReturnError(dbError)

		// Call the repository function
		_, err := repo.GetByID(3)

		// Assertions
		assert.Error(t, err)
		assert.Equal(t, dbError, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
