package model

import (
	"github.com/BioChemML/SIC50/server/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	if config.Cfg.SqlitePath != "" {
		db, err = getSqliteDB(config.Cfg.SqlitePath)
	} else {
		db, err = getMysqlDB()
	}
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Task{})
}

func getSqliteDB(p string) (*gorm.DB, error) {
	opts := &gorm.Config{
		SkipDefaultTransaction: true, PrepareStmt: true,
	}
	return gorm.Open(sqlite.Open(p), opts)
}

func getMysqlDB() (*gorm.DB, error) {
	return nil, nil
}
