// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	movies "movies/definitions/movies"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddToFavorites provides a mock function with given fields: movieID
func (_m *Service) AddToFavorites(movieID int) error {
	ret := _m.Called(movieID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(movieID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CSVImport provides a mock function with given fields:
func (_m *Service) CSVImport() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: request
func (_m *Service) Create(request movies.CreateMovieRequest) (movies.Movie, error) {
	ret := _m.Called(request)

	var r0 movies.Movie
	if rf, ok := ret.Get(0).(func(movies.CreateMovieRequest) movies.Movie); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(movies.Movie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(movies.CreateMovieRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *Service) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *Service) FindAll(ctx *gin.Context) movies.PaginatedMovies {
	ret := _m.Called(ctx)

	var r0 movies.PaginatedMovies
	if rf, ok := ret.Get(0).(func(*gin.Context) movies.PaginatedMovies); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(movies.PaginatedMovies)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *Service) GetByID(id int) (movies.Movie, error) {
	ret := _m.Called(id)

	var r0 movies.Movie
	if rf, ok := ret.Get(0).(func(int) movies.Movie); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(movies.Movie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExtendedMovieInfoByID provides a mock function with given fields: request
func (_m *Service) GetExtendedMovieInfoByID(request movies.ExtendMovieInfoRequest) (string, error) {
	ret := _m.Called(request)

	var r0 string
	if rf, ok := ret.Get(0).(func(movies.ExtendMovieInfoRequest) string); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(movies.ExtendMovieInfoRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchTMDBByNameAndYear provides a mock function with given fields: request
func (_m *Service) SearchTMDBByNameAndYear(request movies.SearchTMDBRequest) (movies.MovieDetailResponse, error) {
	ret := _m.Called(request)

	var r0 movies.MovieDetailResponse
	if rf, ok := ret.Get(0).(func(movies.SearchTMDBRequest) movies.MovieDetailResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(movies.MovieDetailResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(movies.SearchTMDBRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, request
func (_m *Service) Update(id int, request movies.UpdateMovieRequest) (movies.Movie, error) {
	ret := _m.Called(id, request)

	var r0 movies.Movie
	if rf, ok := ret.Get(0).(func(int, movies.UpdateMovieRequest) movies.Movie); ok {
		r0 = rf(id, request)
	} else {
		r0 = ret.Get(0).(movies.Movie)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, movies.UpdateMovieRequest) error); ok {
		r1 = rf(id, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
