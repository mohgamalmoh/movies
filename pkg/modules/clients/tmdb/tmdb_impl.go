package tmdb

import (
	"encoding/json"
	"fmt"
	"io"
	"movies/definitions/movies"
	"net/http"
	"net/url"
)

type Client struct {
}

func NewClient() Client {
	return Client{}
}

func (c Client) SearchTMDBByNameAndYear(request movies.SearchTMDBRequest) (movies.MovieDetailResponse, error) {
	encodedQuery := url.QueryEscape(request.Query)
	url := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?query=%s&year=%s", encodedQuery, request.Year)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MjI5YWJiNjg2YWI4MmIzOTM5NjhmNzY5ODlkZmJhYSIsIm5iZiI6MTcyNDMxNDgyNy43NzQyMDcsInN1YiI6IjVlNTJlNWZmZDJjMGMxMDAxN2EzMjQ4ZSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.0Dy0nMhmzR8uDwnsjBU1d8KXH_KKUW5hHB1daVvYlCg")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var moviesResponse movies.MovieDetailResponse
	err := json.Unmarshal(body, &moviesResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return movies.MovieDetailResponse{}, err
	}
	return moviesResponse, nil
}
