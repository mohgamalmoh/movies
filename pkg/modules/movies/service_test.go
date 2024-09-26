package movies_test

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/stretchr/testify/assert"
	moviesDefinitions "movies/definitions/movies"
	authMocks "movies/mocks/definitions/auth"
	redisMocks "movies/mocks/definitions/clients/redis"
	TMDBMocks "movies/mocks/definitions/clients/tmdb"
	mocks "movies/mocks/definitions/movies"
	"movies/pkg/modules/movies"
	"strconv"
	"testing"
)

func setupService() (moviesDefinitions.Service, *mocks.Repository) {

	tmdbClient := new(TMDBMocks.TMDBClient)
	authClient := new(authMocks.AuthService)
	redisClient := new(redisMocks.RedisClient)
	repo := new(mocks.Repository)
	return movies.NewService(tmdbClient, authClient, redisClient, repo), repo
}

func TestService_Create(t *testing.T) {
	t.Run(
		"create movie - success", func(t *testing.T) {
			// setup
			service, repoMock := setupService()

			// Assert the input request for creating a movie
			createMovieRequest := moviesDefinitions.CreateMovieRequest{
				Name:     "Inception",
				Genre:    "Sci-Fi",
				Overview: "A mind-bending thriller",
				Year:     "2010",
			}

			// Expected Movie object to be created
			movie := moviesDefinitions.Movie{
				Name:     createMovieRequest.Name,
				Genre:    createMovieRequest.Genre,
				Overview: createMovieRequest.Overview,
				Year:     createMovieRequest.Year,
			}

			// Expected response for the service call
			expectedMovie := moviesDefinitions.Movie{
				Id:       strconv.Itoa(1),
				Name:     movie.Name,
				Genre:    movie.Genre,
				Overview: movie.Overview,
				Year:     movie.Year,
			}

			// Mock the repository's Create method to return the created movie without error
			repoMock.On("Create", movie).Return(expectedMovie, nil)

			// Call the service's Create method
			actualResponse, err := service.Create(createMovieRequest)

			// Assertions
			assert.NoError(t, err)
			assert.Equal(t, expectedMovie, actualResponse)

			// Verify that the expectations set on the mock repository are met
			repoMock.AssertExpectations(t)
		},
	)

	t.Run(
		"create movie - db error", func(t *testing.T) {
			// setup
			service, repoMock := setupService()

			// Assert the input request for creating a movie
			createMovieRequest := moviesDefinitions.CreateMovieRequest{
				Name:     "Inception",
				Genre:    "Sci-Fi",
				Overview: "A mind-bending thriller",
				Year:     "2010",
			}

			// Expected Movie object to be created
			movie := moviesDefinitions.Movie{
				Name:     createMovieRequest.Name,
				Genre:    createMovieRequest.Genre,
				Overview: createMovieRequest.Overview,
				Year:     createMovieRequest.Year,
			}

			// Mock the repository's Create method to return the created movie without error
			repoMock.On("Create", movie).Return(moviesDefinitions.Movie{}, errors.New("DB error"))

			// Call the service's Create method
			actualResponse, err := service.Create(createMovieRequest)

			// Assertions
			assert.Error(t, err)
			assert.Equal(t, actualResponse, moviesDefinitions.Movie{})
			assert.Equal(t, fmt.Errorf("failed to create movie: %v", errors.New("DB error")), err)

			// Verify that the expectations set on the mock repository are met
			repoMock.AssertExpectations(t)
		},
	)
}

func TestService_FindAll(t *testing.T) {
	t.Run(
		"list all - success", func(t *testing.T) {
			// setup
			service, repoMock := setupService()

			assertedContext := gin.Context{}

			assertedMovies := []moviesDefinitions.Movie{
				{
					Id:    strconv.Itoa(1),
					Name:  "The Shawshank Redemption",
					Genre: "thriller",
				},
				{
					Id:    strconv.Itoa(2),
					Name:  "The Shawshank Call",
					Genre: "comedy",
				},
			}

			assertedPagination := paginate.Page{
				Page:       1,
				Size:       2,
				TotalPages: 100,
				Total:      200,
			}

			expectedMovies := []moviesDefinitions.MovieResponse{
				{
					Id:    assertedMovies[0].Id,
					Name:  assertedMovies[0].Name,
					Genre: assertedMovies[0].Genre,
				},
				{
					Id:    assertedMovies[1].Id,
					Name:  assertedMovies[1].Name,
					Genre: assertedMovies[1].Genre,
				},
			}

			expectedPagination := moviesDefinitions.Pagination{
				CurrentPage: assertedPagination.Page,
				PerPage:     assertedPagination.Size,
				TotalPages:  assertedPagination.TotalPages,
				Total:       assertedPagination.Total,
			}

			expectedResponse := moviesDefinitions.PaginatedMovies{
				Movies:     expectedMovies,
				Pagination: expectedPagination,
			}

			// mocks
			repoMock.On("FindAll", &assertedContext).Return(assertedPagination, assertedMovies)

			// call method
			actualResponse := service.FindAll(&assertedContext)

			assert.Equal(t, expectedResponse, actualResponse)
			repoMock.AssertExpectations(t)
		},
	)
}
