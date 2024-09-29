package boot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	log "github.com/sirupsen/logrus"
	"movies/config"
	"movies/pkg/database"
	"movies/pkg/modules/auth"
	redis2 "movies/pkg/modules/clients/redis"
	"movies/pkg/modules/clients/tmdb"
	"movies/pkg/modules/helper"
	"movies/pkg/modules/movies"
	"movies/router"
)

func APIBoot(cfg config.API) *gin.Engine {

	fmt.Println("starting API ......")
	db, err := database.NewDatabase(cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	pg := paginate.New()
	// Repository
	moviesRepository := movies.NewRepository(db, pg)

	//clients
	tMDBClient := tmdb.NewClient()

	redisConn := config.RedisConnection(cfg.Redis)
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
