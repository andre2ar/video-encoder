package database

import (
	"log"

	"github.com/andre2ar/video-encoder/domain"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db          *gorm.DB
	Dsn         string
	DsnTest     string
	DbType      string
	DbTypeTest  string
	Debug       bool
	AutoMigrate bool
	Env         string
}

func NewDatabase() *Database {
	return &Database{}
}

func NewDatabaseTest() *gorm.DB {
	dbInstance := NewDatabase()

	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory"
	dbInstance.AutoMigrate = true
	dbInstance.Debug = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Test DB error: %v", err)
	}

	return connection
}

func (d *Database) getDialector(dbType, dsn string) gorm.Dialector {
	switch dbType {
	case "postgres", "postgresql":
		return postgres.Open(dsn)
	case "sqlite", "sqlite3":
		return sqlite.Open(dsn)
	default:
		log.Fatalf("Unsupported database type: %s", dbType)
		return nil
	}
}

func (d *Database) Connect() (*gorm.DB, error) {
	var dialector gorm.Dialector
	var err error

	config := &gorm.Config{}
	if d.Debug {
		config.Logger = logger.Default.LogMode(logger.Info)
	}

	if d.Env == "test" {
		dialector = d.getDialector(d.DbTypeTest, d.DsnTest)
	} else {
		dialector = d.getDialector(d.DbType, d.Dsn)
	}

	d.Db, err = gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	if d.AutoMigrate {
		err := d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		if err != nil {
			return nil, err
		}
	}

	return d.Db, nil
}
