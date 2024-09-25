package movies

type ListMoviesRequest struct {
	Name    *string `validate:"required,min=1,max=200" form:"name"`
	Genre   *string `validate:"required,min=1,max=200" form:"genre"`
	Page    int     `form:"page"`
	PerPage int     `form:"per_page"`
}

type AddToFavouritesRequest struct {
	MovieID string `validate:"required,min=1,max=200" json:"movie_id"`
}

type ExtendMovieInfoRequest struct {
	ID int `validate:"required" form:"id"`
}

type SearchTMDBRequest struct {
	Query string `validate:"required,min=1,max=200" form:"query"`
	Year  string `validate:"required,min=1,max=200" form:"year"`
}

type UpdateMovieRequest struct {
	Id       int    `validate:"required"`
	Overview string `json:"overview"`
}

type MovieDetailResponse struct {
	Data []MovieOverview `json:"results"`
}

type MovieOverview struct {
	Overview string `json:"overview"`
}

type MovieResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
}

type Pagination struct {
	CurrentPage int64 `json:"page"`
	PerPage     int64 `json:"size"`
	TotalPages  int64 `json:"total_pages"`
	Total       int64 `json:"total"`
}

type PaginatedMovies struct {
	Movies     []MovieResponse
	Pagination Pagination
}

type ListingResponse struct {
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}
