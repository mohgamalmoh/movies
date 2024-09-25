package migrations

import (
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	db "gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

const (
	DefaultMigrationsFolder = "migrations/sql/"
)

func Run(db db.DB, migrationFolder ...string) {
	flag.Parse()
	run(db, flag.Args(), migrationFolder...)
}

func RunWithArgs(db db.DB, args []string, migrationFolder ...string) {
	run(db, args, migrationFolder...)
}

func run(db db.DB, args []string, migrationFolder ...string) {
	var err error

	migrationPath := DefaultMigrationsFolder
	if len(migrationFolder) > 0 {
		migrationPath = migrationFolder[0]
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("error getting sql.DB representation:", err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal("error instantiating postgres instance:", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://"+migrationPath, "postgres", driver)
	if err != nil {
		log.Fatal("error instantiating migration instance:", err)
	}

	startTime := time.Now()

	switch args[0] {
	case "up":
		err = m.Up()
	case "down":
		err = m.Steps(-1)
	case "clear":
		err = m.Down()
	case "force":
		version, _ := strconv.Atoi(args[1])
		err = m.Force(version)
	case "create":
		args := args[1:]
		createFlagSet := flag.NewFlagSet("create", flag.ExitOnError)
		_ = createFlagSet.Parse(args)

		if createFlagSet.NArg() == 0 {
			log.Fatal("Specify a name for the migration")
		}

		createCmd(migrationPath, startTime.Unix(), createFlagSet.Arg(0))
	}

	if err != nil {
		if err != migrate.ErrNoChange {
			log.Fatal(err)
		} else {
			log.Info(err)
		}
	}

	log.Info("Finished after: ", time.Since(startTime).String())
}

func createCmd(migrationPath string, timestamp int64, name string) {
	base := fmt.Sprintf("%v%v_%v.", migrationPath, timestamp, name)
	_ = os.MkdirAll(migrationPath, os.ModePerm)
	createFile(base + "up.sql")
	createFile(base + "down.sql")
}

func createFile(fname string) {
	if _, err := os.Create(fname); err != nil {
		log.Fatal(err)
	}
}
