package database

import (
	"fmt"
	"github.com/alexandrebrundias/product-crud/core"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type Database struct {
	db          *gorm.DB
	dsn         string
	dialect     string
	debug       bool
	autoMigrate bool
	env         string
}

func NewDatabaseTest() *Database {
	return &Database{dsn: ":memory:", dialect: "sqlite3", debug: true, autoMigrate: true}
}

func NewDatabase() *Database {
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	name := viper.GetString("database.name")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")
	debug := viper.GetBool("database.config.debug")
	autoMigrate := viper.GetBool("database.config.auto-migrate")
	dialect := viper.GetString("database.config.dialect")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)

	return &Database{dsn: dsn, dialect: dialect, debug: debug, autoMigrate: autoMigrate}
}

func (d Database) Connect() (*gorm.DB, error) {
	var err error
	d.db, err = gorm.Open(d.dialect, d.dsn)

	if err != nil {
		return nil, err
	}

	d.db.SingularTable(true)
	d.db.LogMode(d.debug)

	if d.autoMigrate {
		d.db.AutoMigrate(&core.Product{})
	}

	return d.db, nil
}
