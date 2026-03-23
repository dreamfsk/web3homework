package config

import (
	"com.dreamfsk/blog/commons"
	"com.dreamfsk/blog/repository/models"
	"com.dreamfsk/blog/utils"
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"path/filepath"
)

const (
	DBTypeSQLite DBTypeConfig = "sqlite"
	DBTypeMySQL  DBTypeConfig = "mysql"
)

func getDBType(c *Config) DBTypeConfig {
	switch c.DBType {
	case "mysql":
		return DBTypeMySQL
	default:
		return DBTypeSQLite
	}
}

// GetDB

func GetDB() (*gorm.DB, error) {
	if DB != nil {
		if commons.Migrate {
			models.Migrate(DB)
		}
		return DB, nil
	}
	err := initDatabase()
	if err != nil {
		return nil, err
	}
	if commons.Migrate {
		models.Migrate(DB)
	}
	return DB, nil
}

// initDatabase /
func initDatabase() error {
	c := Env()
	dbType := getDBType(c)
	switch dbType {
	case DBTypeSQLite:
		err := initSqlite(c)
		if err != nil {
			return err
		}
	case DBTypeMySQL:
		err := initMysql(c)
		if err != nil {
			return err
		}
	default:
		log.Printf("unsupported database type: %s", dbType)
	}
	return nil
}

func initMysql(c *Config) error {
	var err error
	dbHost := c.Mysql.Host
	dbPort := c.Mysql.Port
	dbUser := c.Mysql.Username
	dbPassword := c.Mysql.Password
	dbName := c.Mysql.DBName
	// host
	link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)
	// connect
	DB, err = gorm.Open(mysql.Open(link), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to MySQL database:", err)
	}
	if DB != nil && c.Mysql.Migrate {
		commons.Migrate = c.Mysql.Migrate
	}
	return err
}

func initSqlite(c *Config) error {
	dbDir, e := utils.GetDbPath()
	if e != nil {
		log.Fatal(" Not Find SqliteL Path:\n", e)
	}
	dbPath := filepath.Join(dbDir, c.Sqlite.DBName)
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		// Logger configuration
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
			NoLowerCase:   false,
			NameReplacer:  nil,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to SqliteL database:", err)
	}
	if DB != nil && c.Mysql.Migrate {
		commons.Migrate = c.Mysql.Migrate
	}
	return err
}
