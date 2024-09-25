package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func NewDatabase() (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	gormConfig := getCfg()
	db, err := gorm.Open(postgres.Open(sqlInfo), &gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getCfg() gorm.Config {
	return gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
}

/*
func NewTestDatabase(
	con gorm.Dialector, maxIdleConnections, maxOpenConnections, connectionMaxLifeTime int,
) (db.Database, error) {
	gCfg := getCfg()
	database, err := commonDB.New(con, gCfg, maxIdleConnections, maxOpenConnections, connectionMaxLifeTime)
	if err != nil {
		return nil, err
	}

	return database.Session(&gorm.Session{}), nil
}
*/
