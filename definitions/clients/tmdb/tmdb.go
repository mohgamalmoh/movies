package tmdb

import (
	"movies/definitions/movies"
)

type TMDBClient interface {
	SearchTMDBByNameAndYear(req movies.SearchTMDBRequest) (movies.MovieDetailResponse, error)
}
