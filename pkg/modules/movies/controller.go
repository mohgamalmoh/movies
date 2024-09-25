package movies

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"movies/definitions/auth"
	"movies/definitions/movies"
	"movies/pkg/modules/helper"
	"net/http"
	"strconv"
)

type Controller struct {
	auth          auth.AuthService
	moviesService movies.Service
	helper        helper.Helper
}

func NewController(auth auth.AuthService, service movies.Service, helper helper.Helper) *Controller {
	return &Controller{
		auth:          auth,
		moviesService: service,
		helper:        helper,
	}
}

// @Summary Get Movies
// @Description Get a paginated list of movies by genre
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Param size query int true "Size of the page"
// @Param genre query string false "Genre of the movies"
// @Success 200 {object} helper.ResponseBody{} "List of movies"
// @Router /movies [get]
func (c *Controller) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll movies")
	moviesResponse := c.moviesService.FindAll(ctx)
	res := movies.ListingResponse{
		Data:       moviesResponse.Movies,
		Pagination: moviesResponse.Pagination,
	}
	c.helper.RespondWithJSON(ctx, http.StatusOK, res)

}

// @Summary Get Extended Moview info
// @Description Get more info about a movie
// @Tags Movies extended info
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Param size query int true "Size of the page"
// @Param genre query string false "Genre of the movies"
// @Success 200 {object} helper.ResponseBody{} "List of movies"
// @Router /movies/info/{id} [get]
func (c *Controller) GetExtendedMovieInfo(ctx *gin.Context) {
	movieID := ctx.Param("id")
	id, err := strconv.Atoi(movieID)
	if err != nil {
		return
	}
	req := movies.ExtendMovieInfoRequest{ID: id}
	movieDetails, err := c.moviesService.GetExtendedMovieInfoByID(req)
	if err != nil {
		return
	}
	c.helper.RespondWithJSON(ctx, http.StatusOK, movieDetails)

}

// @Summary Add Movie to Favorites
// @Description Add Movie to Favorites
// @Tags Favorite Movies
// @Accept  json
// @Produce  json
// @Param movie_id body int true "movie_id"
// @Success 200 {object} helper.ResponseBody{response.MovieDetailResponse{}} "movie extended info from TMDB API"
// @Router /movies/favorite [post]
func (c *Controller) AddToFavorites(ctx *gin.Context) {
	req := movies.AddToFavouritesRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, nil)
	}
	id, err := strconv.Atoi(req.MovieID)
	err = c.moviesService.AddToFavorites(id)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, nil)
	} else {
		c.helper.RespondWithJSON(ctx, http.StatusOK, nil)
	}
}

// @Summary CSV import
// @Description import data from google sheet
// @Tags import
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.ResponseBody{}
// @Router /movies/import [get]
func (c *Controller) CSVImport(ctx *gin.Context) {
	err := c.moviesService.CSVImport()
	if err != nil {
		return
	}

	c.helper.RespondWithJSON(ctx, http.StatusOK, nil)

}
