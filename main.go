package main

import (
	"fmt"
	"log"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/mtrentz/dados-cnpj/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	// Connect to database, set it to global variable
	fmt.Println("Connecting to database...")
	var err error
	dsn := "root:999999@tcp(127.0.0.1:3366)/dados_cnpj?charset=latin2&collation=latin2_general_ci&autocommit=false&parseTime=true"
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Unable to open connection to db: ", err)
	}

	// dataDir := "data"
	// OrganizedFiles := organize.GetFiles(dataDir)

	// transform.Unzip(OrganizedFiles, dataDir)

	// transform.ConcatAll(dataDir)

	database.InsertAll()
	// database.TestInsert()

	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y0.D10911.SOCIOCSV", false)
	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y1.D10911.SOCIOCSV", false)

	// transform.GetFilesByCategory(dataDir)
	// transform.ConcatCsvs()
}
