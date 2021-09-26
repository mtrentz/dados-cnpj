package main

import (
	"github.com/mtrentz/dados-cnpj/transform"
)

func main() {
	dataDir := "data"
	// OrganizedFiles := organize.GetFiles(dataDir)
	// for i, e := range OrganizedFiles {
	// 	fmt.Println(i, ":")
	// 	for _, val := range e {
	// 		fmt.Println(">>>", val)
	// 	}
	// }

	// transform.Unzip(OrganizedFiles, dataDir)

	transform.ConcatAll(dataDir)

	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y0.D10911.SOCIOCSV", false)
	// transform.AppendAllLines("data/socio/socios.csv", "data/socio/K3241.K03200Y1.D10911.SOCIOCSV", false)

	// transform.GetFilesByCategory(dataDir)
	// transform.ConcatCsvs()
}
