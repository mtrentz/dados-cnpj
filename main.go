package main

import "github.com/mtrentz/dados-cnpj/transform"

func main() {
	// dataDir := "data"
	// 	OrganizedFiles := organize.GetFiles(dataDir)
	// 	// transform.Unzip()
	// 	transform.Unzip(&OrganizedFiles, dataDir)

	// transform.GetFilesByCategory(dataDir)
	transform.ConcatCsvs()
}
