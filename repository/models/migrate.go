package models

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {

	// autoMigrate
	//err := db.AutoMigrate(&User{}, &Post{}, &Comment{},&SysError{})
	err := db.AutoMigrate(&SysError{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("===migrated tables successfully===")

}
