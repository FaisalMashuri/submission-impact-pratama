package driver

import (
	"fmt"

	"github.com/FaisalMashuri/submission-golang/app/entity"
	"github.com/FaisalMashuri/submission-golang/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatababseConfiguration struct {
	config *config.Config
}

func InitDB(DBUSER, DBPASSWORD, HOSTDB, DBPORT, DBNAME string) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUSER, DBPASSWORD, HOSTDB, DBPORT, DBNAME)
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&entity.ProductEntity{})
	return DB
}
