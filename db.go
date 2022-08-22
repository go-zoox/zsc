package zsc

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	if db == nil {
		panic("DB is not initialized")
	}

	return db
}

func LoadDB(engine string, dsn string) (err error) {
	var dialector gorm.Dialector
	switch engine {
	case "postgres":
		dialector = postgres.Open(dsn)
	case "mysql":
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		panic(fmt.Errorf("unknown engine: %s", engine))
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:               logger.Default.LogMode(logger.Info), // Print SQL queries
		DisableAutomaticPing: false,
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return fmt.Errorf("connecting database failed: %s", err.Error())
	}

	return nil
}
