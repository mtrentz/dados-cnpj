package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Connect to database, set it to global variable
	fmt.Println("Connecting to database...")

	var err error

	dsn := "dadoscnpj:dadoscnpj@tcp(dados-cnpj_db:3306)/dados_cnpj?charset=latin2&collation=latin2_general_ci&autocommit=false&parseTime=true"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Unable to open connection to db: ", err)
	}

}
