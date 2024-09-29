package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"movies/config"
	"movies/pkg/database"
	"movies/pkg/migrations"
)

func Migrations() *cobra.Command {
	return &cobra.Command{
		Use:   "migrations",
		Short: "Run app database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Starting app migrations")
			runMigrations(args)
		},
	}
}

func runMigrations(args []string) {
	cfg, err := config.NewMigrationsConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}
	migrations.RunWithArgs(*db, args)
}
