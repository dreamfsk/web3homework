package config

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"homework03/models"
	"homework03/utils"
	"log"
	"path/filepath"
)

// export db
var DB *gorm.DB

type DBType string

const (
	DBTypeSQLite DBType = "sqlite"
	DBTypeMySQL  DBType = "mysql"
)

func getDBType() DBType {
	utils.LoadEnv()
	dbType := utils.GetEnv("DB_TYPE", "DB_TYPE")
	switch dbType {
	case "mysql":
		return DBTypeMySQL
	default:
		return DBTypeSQLite
	}
}

// GetDB

func GetDB(migrate bool) *gorm.DB {
	if DB != nil {
		return DB
	}
	initDatabase(migrate)
	return DB
}

// initDatabase /
func initDatabase(migrate bool) {
	dbType := getDBType()
	switch dbType {
	case DBTypeSQLite:
		initSqlite()
	case DBTypeMySQL:
		initMysql()
	default:
		log.Printf("unsupported database type: %s", dbType)
	}
	if DB != nil && migrate {
		initModels()
	}
}

func initMysql() {
	var err error
	dbHost := utils.GetEnv("DB_HOST", "localhost")
	dbPort := utils.GetEnv("DB_PORT", "3306")
	dbUser := utils.GetEnv("DB_USER", "root")
	dbPassword := utils.GetEnv("DB_PASSWORD", "")
	dbName := utils.GetEnv("DB_NAME", "golang_blog")
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
}

func initSqlite() {

	dbDir, e := utils.GetPath()
	if e != nil {
		log.Fatal(" Not Find SqliteL Path:\n", e)
	}
	dbName := utils.GetEnv("SQ_DB", "blog.sqlite.db")
	dbPath := filepath.Join(dbDir, dbName)
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
}

func initModels() {
	var err error
	if DB == nil {
		log.Println("DB is null")
		return
	}
	// autoMigrate
	err = DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("===database connected and migrated successfully===")
}
