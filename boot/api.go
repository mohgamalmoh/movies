package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	log "github.com/sirupsen/logrus"
	"movies/config"
	movies2 "movies/definitions/movies"
	"movies/definitions/movies_sync_status"
	"movies/definitions/users"
	"movies/pkg/database"
	"movies/pkg/modules/auth"
	redis2 "movies/pkg/modules/clients/redis"
	"movies/pkg/modules/clients/tmdb"
	"movies/pkg/modules/helper"
	"movies/pkg/modules/movies"
	"movies/router"
)

const ServiceName = "ads-auctions"

func APIBoot() *gin.Engine {

	fmt.Println("starting API ......")
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.Table("movies").AutoMigrate(&movies2.Movie{})
	db.Table("users").AutoMigrate(&users.User{})
	db.Table("users_movies").AutoMigrate(&users.UsersMovies{})
	db.Table("movies_sync_status").AutoMigrate(&movies_sync_status.MoviesSyncStatus{})

	pg := paginate.New()
	// Repository
	moviesRepository := movies.NewRepository(db, pg)

	//clients
	tMDBClient := tmdb.NewClient()

	redisConn := config.RedisConnection()
	redisClient := redis2.NewRedisClient(redisConn)
	// Service
	authSerivce := auth.NewAuthServiceImpl()
	moviesService := movies.NewService(tMDBClient, authSerivce, redisClient, moviesRepository)

	//helper
	helpr := helper.NewHelper()
	// Controller
	moviesController := movies.NewController(authSerivce, moviesService, helpr)

	// Router
	routes := router.NewRouter(moviesController)

	return routes
}
