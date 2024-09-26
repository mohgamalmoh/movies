package helper

import (
	"github.com/gin-gonic/gin"
	"movies/definitions/movies"
)

type Helper struct{}

func NewHelper() Helper {
	return Helper{}
}

func (h Helper) ErrorPanic(err error) {
	if err != nil {
		panic(err)
	}
}

type ResponseBody struct {
	Body interface{} `json:"response"`
}

func (h Helper) RespondWithJSON(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, ResponseBody{Body: data})
}

func (h Helper) MapMovieToMovieResponse(movie movies.Movie) movies.MovieResponse {
	return movies.MovieResponse{
		Id:       movie.Id,
		Name:     movie.Name,
		Genre:    movie.Genre,
		Year:     movie.Year,
		Overview: movie.Overview,
	}
}
