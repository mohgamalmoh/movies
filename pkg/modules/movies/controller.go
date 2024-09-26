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

// @Summary Create a new movie
// @Description Add a new movie to the database
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param movie body movies.CreateMovieRequest true "Movie Data"
// @Success 201 {object} helper.ResponseBody{response.MovieDetailResponse{}} "Movie created successfully"
// @Failure 400 {object} helper.ResponseBody{} "Invalid request data"
// @Router /movies [post]
func (c *Controller) Create(ctx *gin.Context) {
	log.Info().Msg("create movie")

	// Bind JSON request data to the CreateMovieRequest struct
	var req movies.CreateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}

	// Call service to create the movie
	movie, err := c.moviesService.Create(req)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusInternalServerError, "Could not create movie")
		return
	}

	c.helper.RespondWithJSON(ctx, http.StatusCreated, c.helper.MapMovieToMovieResponse(movie))
}

// @Summary Get movie by ID
// @Description Get a specific movie by its ID
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 200 {object} helper.ResponseBody{response.MovieDetailResponse{}} "Movie details"
// @Failure 404 {object} helper.ResponseBody{} "Movie not found"
// @Router /movies/{id} [get]
func (c *Controller) GetByID(ctx *gin.Context) {
	log.Info().Msg("get movie by ID")

	// Parse movie ID from the URL parameter
	movieID := ctx.Param("id")
	id, err := strconv.Atoi(movieID)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, "Invalid movie ID")
		return
	}

	// Call service to get movie details by ID
	movie, err := c.moviesService.GetByID(id)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusNotFound, "Movie not found")
		return
	}

	// Return movie details
	c.helper.RespondWithJSON(ctx, http.StatusOK, movie)
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

// @Summary Update a movie
// @Description Update details of an existing movie
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Param movie body movies.UpdateMovieRequest true "Updated Movie Data"
// @Success 200 {object} helper.ResponseBody{response.MovieDetailResponse{}} "Movie updated successfully"
// @Failure 400 {object} helper.ResponseBody{} "Invalid request data"
// @Router /movies/{id} [put]
func (c *Controller) Update(ctx *gin.Context) {
	log.Info().Msg("update movie")

	// Parse movie ID from the URL parameter
	movieID := ctx.Param("id")
	id, err := strconv.Atoi(movieID)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, "Invalid movie ID")
		return
	}

	// Bind JSON request data to the UpdateMovieRequest struct
	var req movies.UpdateMovieRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, "Invalid request data")
		return
	}

	// Call service to update the movie
	movie, err := c.moviesService.Update(id, req)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusInternalServerError, "Could not update movie")
		return
	}

	// Return success response with updated movie details
	c.helper.RespondWithJSON(ctx, http.StatusOK, movie)
}

// @Summary Delete a movie
// @Description Delete a movie by its ID
// @Tags Movies
// @Accept  json
// @Produce  json
// @Param id path int true "Movie ID"
// @Success 204 {object} helper.ResponseBody{} "Movie deleted successfully"
// @Failure 404 {object} helper.ResponseBody{} "Movie not found"
// @Router /movies/{id} [delete]
func (c *Controller) Delete(ctx *gin.Context) {
	log.Info().Msg("delete movie")

	// Parse movie ID from the URL parameter
	movieID := ctx.Param("id")
	id, err := strconv.Atoi(movieID)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusBadRequest, "Invalid movie ID")
		return
	}

	// Call service to delete the movie
	err = c.moviesService.Delete(id)
	if err != nil {
		c.helper.RespondWithJSON(ctx, http.StatusNotFound, "Movie not found")
		return
	}

	// Return no content status on successful deletion
	c.helper.RespondWithJSON(ctx, http.StatusNoContent, nil)
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
