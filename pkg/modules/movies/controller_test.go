package movies_test

import (
	"errors"
	"fmt"
	moviesDef "movies/definitions/movies"
	moviesMock "movies/mocks/definitions/movies"
	"movies/pkg/modules/helper"
	"movies/pkg/modules/movies"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestController_GetByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Initialize the mock service and helper
	mockService := new(moviesMock.Service)
	mockHelper := new(helper.Helper)

	// Create a new controller with the mock service and helper
	controller := movies.NewController(nil, mockService, *mockHelper)

	// Define a new gin router
	router := gin.Default()
	router.GET("/movies/:id", controller.GetByID)

	// Define a mock movie for testing
	expectedMovie := moviesDef.Movie{
		Id:       "1",
		Name:     "The Shawshank Redemption",
		Genre:    "Drama",
		Year:     "1994",
		Overview: "Two imprisoned men bond over a number of years...",
	}

	t.Run("Success - Get movie by ID", func(t *testing.T) {
		// Mock the service layer
		mockService.On("GetByID", 1).Return(expectedMovie, nil)

		// Create an HTTP request for /moviesDef/1
		req, _ := http.NewRequest(http.MethodGet, "/movies/1", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assert the status code
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert the response contains the expected movie details
		expectedResponse := fmt.Sprintf(`{"Id":"%s","Name":"%s","Genre":"%s","Year":"%s","Overview":"%s"}`,
			expectedMovie.Id, expectedMovie.Name, expectedMovie.Genre, expectedMovie.Year, expectedMovie.Overview)
		assert.Contains(t, w.Body.String(), expectedResponse)

		mockService.AssertExpectations(t)
	})

	t.Run("Failure - Movie not found", func(t *testing.T) {
		// Mock service behavior for a non-existent movie
		mockService.On("GetByID", 999).Return(moviesDef.Movie{}, errors.New("movie not found"))

		// Create an HTTP request for /moviesDef/999
		req, _ := http.NewRequest(http.MethodGet, "/movies/999", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assert the status code is 404 Not Found
		assert.Equal(t, http.StatusNotFound, w.Code)

		// Assert the response contains the correct error message
		assert.Contains(t, w.Body.String(), "Movie not found")

		mockService.AssertExpectations(t)
	})

	t.Run("Failure - Invalid movie ID", func(t *testing.T) {
		// Create an HTTP request with an invalid movie ID
		req, _ := http.NewRequest(http.MethodGet, "/movies/invalid-id", nil)
		w := httptest.NewRecorder()

		// Perform the request
		router.ServeHTTP(w, req)

		// Assert the status code is 400 Bad Request
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Assert the response contains the correct error message
		assert.Contains(t, w.Body.String(), "Invalid movie ID")
	})
}
