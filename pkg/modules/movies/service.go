package movies

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"io"
	"movies/definitions/auth"
	redis2 "movies/definitions/clients/redis"
	"movies/definitions/clients/tmdb"
	"movies/definitions/movies"
	moviesDefinition "movies/definitions/movies"
	"net/http"
	"time"
)

type Service struct {
	client           tmdb.TMDBClient
	authService      auth.AuthService
	RedisClient      redis2.RedisClient
	MoviesRepository movies.Repository
}

func NewService(client tmdb.TMDBClient, authService auth.AuthService, redisClient redis2.RedisClient, moviesRepository movies.Repository) movies.Service {
	return &Service{
		client:           client,
		authService:      authService,
		RedisClient:      redisClient,
		MoviesRepository: moviesRepository,
	}
}

// FindAll implements Service
func (s *Service) FindAll(ctx *gin.Context) moviesDefinition.PaginatedMovies {
	pag, items := s.MoviesRepository.FindAll(ctx)

	var movies []moviesDefinition.MovieResponse
	for _, value := range items {
		movie := moviesDefinition.MovieResponse{
			Id:    value.Id,
			Name:  value.Name,
			Genre: value.Genre,
		}
		movies = append(movies, movie)
	}

	pagination := moviesDefinition.Pagination{
		CurrentPage: pag.Page,
		PerPage:     pag.Size,
		TotalPages:  pag.TotalPages,
		Total:       pag.Total,
	}
	PaginateMovies := moviesDefinition.PaginatedMovies{Movies: movies, Pagination: pagination}

	return PaginateMovies
}

func (s *Service) GetByID(id int) movies.Movie {
	movie := s.MoviesRepository.GetByID(id)
	return movie
}

func (s *Service) Update(request moviesDefinition.UpdateMovieRequest) error {
	movie := s.MoviesRepository.Update(request)
	return movie
}

func (s *Service) SearchTMDBByNameAndYear(request moviesDefinition.SearchTMDBRequest) (moviesDefinition.MovieDetailResponse, error) {
	info, err := s.client.SearchTMDBByNameAndYear(request)
	if err != nil {
		return moviesDefinition.MovieDetailResponse{}, nil
	}
	return info, nil
}

func (s *Service) GetExtendedMovieInfoByID(req moviesDefinition.ExtendMovieInfoRequest) (overview string, err error) {

	overview, redisErr := s.RedisClient.GetCache(fmt.Sprintf("movie_overview:%d", req.ID))
	if errors.Is(redisErr, redis.Nil) {
		fmt.Println("Cache miss: Details not found in Redis")
		movie := s.GetByID(req.ID)
		if movie.Overview != "" {
			overview = movie.Overview
		} else {
			fmt.Println("DB miss: Details not found in DB")
			searchRequest := moviesDefinition.SearchTMDBRequest{
				Query: movie.Name,
				Year:  movie.Year,
			}

			movieDetails, err := s.client.SearchTMDBByNameAndYear(searchRequest)
			if err != nil {
				return "", err
			}
			fmt.Print(movieDetails)
			overview = movieDetails.Data[0].Overview

			err = s.RedisClient.SetCache(fmt.Sprintf("movie_overview:%d", req.ID), overview, 1*time.Minute)
			if err != nil {
				return "", err
			}

			err = s.Update(moviesDefinition.UpdateMovieRequest{
				Id:       req.ID,
				Overview: overview,
			})
			if err != nil {
				return "", err
			}
			return overview, nil
		}

	} else if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return
}

func (s *Service) AddToFavorites(movieID int) error {
	movie := s.GetByID(movieID)

	infoReq := movies.SearchTMDBRequest{
		Query: movie.Name,
		Year:  movie.Year,
	}
	movieDetails, err := s.SearchTMDBByNameAndYear(infoReq)
	if err != nil {
		return err
	}

	if len(movieDetails.Data) == 0 {
		return errors.New("No data found in TMDB")
	}

	err = s.Update(movies.UpdateMovieRequest{
		Id:       movieID,
		Overview: movieDetails.Data[0].Overview,
	})
	if err != nil {
		return err
	}
	user := s.authService.GetAuthUser()
	err = s.MoviesRepository.AddToFavorites(user.Id, movieID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) CSVImport() (err error) {
	// Fetch the CSV from Google Drive
	csvBody, err := fetchCSVFromGoogleDrive(googleDriveFileID)
	if err != nil {
		fmt.Printf("Error fetching CSV: %v\n", err)
		return
	}
	defer csvBody.Close()

	lastInsertedMovie, err := s.MoviesRepository.GetLastSyncedMovie()
	if err != nil {
		fmt.Printf("Error getting last synced TMDBID: %v\n", err)
		return
	}

	// Parse the CSV into Movie structs
	movies, err := parseCSV(csvBody, lastInsertedMovie)
	if err != nil {
		fmt.Printf("Error parsing CSV: %v\n", err)
		return
	}

	// Perform bulk insert using batch processing
	if err := s.bulkInsertMovies(movies); err != nil {
		fmt.Printf("Error inserting movies: %v\n", err)
	} else {
		fmt.Println("Movie inserted successfully!")
	}
	return
}

const (
	googleDriveFileID = "1lGwK58azLGL9jETQokdiqbkUlKuiiwOAbMVPtg9IrdM" // Replace with your Google Drive File ID
	batchSize         = 50                                             // Define batch size for bulk insert
)

// getGoogleDriveDirectDownloadLink converts a Google Drive shareable link to a direct download link
func getGoogleDriveDirectDownloadLink(fileID string) string {
	return fmt.Sprintf("https://docs.google.com/spreadsheets/d/%s/export?format=csv", fileID)
}

// fetchCSVFromGoogleDrive fetches the CSV file from Google Drive and returns it as an io.Reader
func fetchCSVFromGoogleDrive(fileID string) (io.ReadCloser, error) {
	url := getGoogleDriveDirectDownloadLink(fileID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CSV: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch CSV: %s", resp.Status)
	}

	return resp.Body, nil
}

// parseCSV parses the CSV from an io.Reader and returns a slice of Movie structs
func parseCSV(reader io.Reader, lastName string) ([]movies.Movie, error) {
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %v", err)
	}

	var movies []movies.Movie
	skip := true
	if lastName == "" {
		skip = false // If there is no last synced movie, import all rows
	}

	for _, record := range records[1:] { // Skip header row
		//releaseYear, _ := strconv.Atoi()
		//rating, _ := strconv.ParseFloat(record[3], 64)

		movie := moviesDefinition.Movie{
			Name:  record[3],
			Genre: record[8],
			Year:  record[5],
		}
		// Skip movies until we reach the last inserted one
		if skip && movie.Name == lastName {
			skip = false // Stop skipping once we reach the last inserted row
			continue
		}

		if !skip {
			movies = append(movies, movie)
		}
	}

	return movies, nil
}

// bulkInsertMovies inserts movies in batches using the specified batch size
func (s *Service) bulkInsertMovies(movies []movies.Movie) error {
	for i := 0; i < len(movies); i += batchSize {
		end := i + batchSize
		if end > len(movies) {
			end = len(movies) // Handle case where the last batch is smaller
		}

		batch := movies[i:end]

		// Insert the batch into the database
		if err := s.MoviesRepository.CreateBatch(&batch); err != nil {
			return fmt.Errorf("failed to insert batch: %v", err)
		}
		lastMovie := batch[len(batch)-1]
		if err := s.MoviesRepository.UpdateSyncStatus(lastMovie.Name); err != nil {
			return fmt.Errorf("failed to update sync status: %v", err)
		}
	}
	return nil
}