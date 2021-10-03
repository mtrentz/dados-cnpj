package database

import (
	"fmt"
	"log"
	"time"

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

	func() {
		// Wait until database is online

		retries := 10
		count := 0

		for {
			DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Silent),
			})
			if err != nil {
				log.Println("Unable to open connection to db: ", err)
			} else {
				// Tudo certo, sai da funcao
				return
			}

			fmt.Println("Trying again in 5 seconds.")
			time.Sleep(time.Second * 5)
			count++

			if count >= retries {
				log.Fatal("Could not connect to database, exiting: ", err)
				return
			}
		}
	}()

}
