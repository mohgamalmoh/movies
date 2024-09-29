package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"movies/config"
)

func NewDatabase(cfg config.Database) (*gorm.DB, error) {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
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
