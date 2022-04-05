package configs

import (
	"projects/entitys"

	"github.com/morkid/paginate"
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

func PaginateConfig() *paginate.Pagination {
	pg := paginate.New(&paginate.Config{
		DefaultSize: 3,
	})
	return pg
}

func AutoMigrate() {
	db := GetConnection()
	db.AutoMigrate(
		entitys.User{},
		entitys.RefreshToken{},
		entitys.Project{},
		entitys.Tasks{},
	)

}
