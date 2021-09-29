package main

import (
	"github.com/mtrentz/dados-cnpj/database"
	"github.com/mtrentz/dados-cnpj/download"
)

func main() {

	database.Connect()

	dataDir := "data"
	download.DownloadAll(dataDir)
	// OrganizedFiles := organize.GetFiles(dataDir)

	// transform.Unzip(OrganizedFiles, dataDir)

	// transform.ConcatAll(dataDir)

	// database.InsertAll()
	// database.TestInsert()

	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y0.D10911.SOCIOCSV", false)
	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y1.D10911.SOCIOCSV", false)

	// transform.GetFilesByCategory(dataDir)
	// transform.ConcatCsvs()
}
