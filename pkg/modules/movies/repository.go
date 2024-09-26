package movies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"movies/definitions/movies"
	"movies/definitions/movies_sync_status"
	"movies/definitions/users"
	"strings"

	"gorm.io/gorm"
)

type Repository struct {
	Db         *gorm.DB
	Pagination *paginate.Pagination
}

func NewRepository(Db *gorm.DB, Pagination *paginate.Pagination) Repository {
	return Repository{Db: Db, Pagination: Pagination}
}

func (t Repository) Create(movie movies.Movie) (movies.Movie, error) {
	err := t.Db.Create(&movie).Error
	if err != nil {
		return movies.Movie{}, err
	}
	return movie, nil
}

func (t Repository) GetByID(id int) (movies.Movie, error) {
	var movie movies.Movie
	err := t.Db.First(&movie, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return movie, fmt.Errorf("movie with ID %d not found", id)
		}
		return movie, err
	}
	return movie, nil
}

func (t Repository) FindAll(ctx *gin.Context) (paginate.Page, []movies.Movie) {
	var movies []movies.Movie
	pg := paginate.New()
	stmt := t.Db.Model(&movies)

	name := strings.ToLower(ctx.Query("name"))
	if name != "" {
		stmt = t.Db.Where("LOWER(name) LIKE ?", name+"%").Model(&movies)
	}

	genre := strings.ToLower(ctx.Query("genre"))
	if genre != "" {
		stmt = t.Db.Where("LOWER(genre) LIKE ?", genre+"%").Model(&movies)
	}

	pagination := pg.With(stmt).Request(ctx.Request).Response(&movies)
	return pagination, movies
}

func (t Repository) Update(movie movies.Movie) error {
	err := t.Db.Save(&movie).Error
	if err != nil {
		return fmt.Errorf("failed to update movie: %v", err)
	}
	return nil
}

func (t Repository) Delete(id int) error {
	var movie movies.Movie
	err := t.Db.First(&movie, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("movie with ID %d not found", id)
		}
		return err
	}

	err = t.Db.Delete(&movie).Error
	if err != nil {
		return fmt.Errorf("failed to delete movie with ID %d: %v", id, err)
	}
	return nil
}

func (t Repository) CreateBatch(movies *[]movies.Movie) error {
	return t.Db.Create(&movies).Error
}

func (t Repository) AddToFavorites(userID int, movieID int) error {
	usersMovies := users.UsersMovies{
		User:  userID,
		Movie: movieID,
	}
	err := t.Db.Create(&usersMovies).Error
	if err != nil {
		return err
	}
	return nil
}

func (t Repository) GetLastSyncedMovie() (string, error) {
	var syncStatus movies_sync_status.MoviesSyncStatus
	if err := t.Db.First(&syncStatus).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil // If no record found, assume no sync has been done
		}
		return "", err
	}
	return syncStatus.Name, nil
}

func (t Repository) UpdateSyncStatus(movieName string) error {
	var syncStatus movies_sync_status.MoviesSyncStatus
	if err := t.Db.First(&syncStatus).Error; err == gorm.ErrRecordNotFound {
		// If no sync status exists, create a new one
		syncStatus = movies_sync_status.MoviesSyncStatus{Name: movieName}
		return t.Db.Create(&syncStatus).Error
	}
	// Update the existing sync status with the new last TMDBID
	syncStatus.Name = movieName
	return t.Db.Save(&syncStatus).Error
}
