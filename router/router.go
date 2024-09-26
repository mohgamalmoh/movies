package router

import (
	"movies/pkg/modules/movies"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(moviesController *movies.Controller) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api/v1")

	moviesRouter := baseRouter.Group("/movies")
	moviesRouter.GET("", moviesController.FindAll)
	moviesRouter.POST("", moviesController.Create)
	moviesRouter.PUT("", moviesController.Update)
	moviesRouter.GET(":id/info/", moviesController.GetExtendedMovieInfo)
	moviesRouter.POST("favorite", moviesController.AddToFavorites)
	moviesRouter.GET("import", moviesController.CSVImport)

	return router
}
