package movies

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	Create(request CreateMovieRequest) (Movie, error)
	FindAll(ctx *gin.Context) PaginatedMovies
	GetByID(id int) (Movie, error)
	Update(id int, request UpdateMovieRequest) (Movie, error)
	SearchTMDBByNameAndYear(request SearchTMDBRequest) (MovieDetailResponse, error)
	GetExtendedMovieInfoByID(request ExtendMovieInfoRequest) (string, error)
	AddToFavorites(movieID int) error
	CSVImport() error
	Delete(id int) error
}
