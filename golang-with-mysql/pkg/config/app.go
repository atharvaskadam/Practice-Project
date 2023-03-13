package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// username: root
// password: ASKASK44
// connection string format: user:pass@tcp(localhost:3306)/dbname?charset=utf8mb4
// root:ASKASK44@tcp(localhost:3306)/book
func Connect() {
	// In order for below code to work, you will have to make sure that you have below server up and running
	 //d, err := gorm.Open("mysql", "root:ASKASK44@tcp(localhost:3306)/book?charset=utf8mb4&parseTime=True&loc=Local")
	d, err := gorm.Open("mysql", "root:ASKASK44@tcp(localhost:3306)/book?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	log.Println("DB Cnnection Successful")
}
func GetDB() *gorm.DB {
	return db
}