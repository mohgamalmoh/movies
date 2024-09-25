package movies_test

import (
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
