package configs

import (
	"projects/entitys"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=dinhnguyen dbname=training port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db

}
func AutoMigrate() {
	db := GetConnection()
	db.AutoMigrate(
		entitys.User{},
		entitys.Images{},
		entitys.RefreshToken{},
		entitys.Project{},
		entitys.Images{},
	)

}
