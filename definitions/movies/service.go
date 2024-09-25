package movies

import (
	"github.com/gin-gonic/gin"
)

type Service interface {
	FindAll(ctx *gin.Context) PaginatedMovies
	GetByID(id int) Movie
	Update(request UpdateMovieRequest) error
	SearchTMDBByNameAndYear(request SearchTMDBRequest) (MovieDetailResponse, error)
	GetExtendedMovieInfoByID(request ExtendMovieInfoRequest) (string, error)
	AddToFavorites(movieID int) error
	CSVImport() error
}
