package cmd

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"movies/boot"
	"movies/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func API() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start app in api mode.",
		Run: func(cmd *cobra.Command, args []string) {

			log.Info("Starting app api")
			runAPI()
		},
	}
}

func runAPI() {

	cfg, err := config.NewAPIConfig()
	if err != nil {
		log.Fatalln(err)
	}
	engine := boot.APIBoot(cfg)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}
}
