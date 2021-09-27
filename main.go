package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/mtrentz/dados-cnpj/organize"
	"github.com/mtrentz/dados-cnpj/transform"
)

func main() {

	// // Connect to database, set it to global variable
	// fmt.Println("Connecting to database...")
	// var err error
	// database.DB, err = sql.Open("mysql", "root:999999@tcp(127.0.0.1:3366)/dados_cnpj?charset=latin2&collation=latin2_general_ci&autocommit=false")
	// if err != nil {
	// 	log.Fatal("Unable to open connection to db: ", err)
	// }
	// defer database.DB.Close()

	dataDir := "data"
	OrganizedFiles := organize.GetFiles(dataDir)

	transform.Unzip(OrganizedFiles, dataDir)

	transform.ConcatAll(dataDir)

	// database.InsertAll()

	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y0.D10911.SOCIOCSV", false)
	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y1.D10911.SOCIOCSV", false)

	// transform.GetFilesByCategory(dataDir)
	// transform.ConcatCsvs()
}
