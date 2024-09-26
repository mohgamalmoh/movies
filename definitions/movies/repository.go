package movies

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
)

type Repository interface {
	Create(movie Movie) (Movie, error)
	Delete(id int) error
	FindAll(ctx *gin.Context) (paginate.Page, []Movie)
	GetByID(id int) (Movie, error)
	Update(Movie) error
	CreateBatch(movies *[]Movie) error
	AddToFavorites(userID int, movieID int) error
	GetLastSyncedMovie() (string, error)
	UpdateSyncStatus(movieName string) error
}
