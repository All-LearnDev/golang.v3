package configs

import (
	"projects/entitys"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "root:dinhnguyen@tcp(127.0.0.1:3309)/gos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
func AutoMigrate() {
	db := GetConnection()
	db.AutoMigrate(&entitys.Book{},
		&entitys.Author{},
		entitys.User{},
		entitys.Images{},
		entitys.JUser{},
		entitys.RefreshToken{},
		entitys.Project{},
		entitys.Developer{},
	)

}
