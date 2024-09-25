package movies

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

type Repository interface {
	FindAll(ctx *gin.Context) (paginate.Page, []Movie)
	GetByID(id int) Movie
	Update(request UpdateMovieRequest) error
	CreateBatch(movies *[]Movie) error
	AddToFavorites(userID int, movieID int) error
	GetLastSyncedMovie() (string, error)
	UpdateSyncStatus(movieName string) error
}
