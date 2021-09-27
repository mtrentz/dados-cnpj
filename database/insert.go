package database

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

func InsertAll() {

	DB.AutoMigrate(&Municipios{})

	csvFile := "data/municipios/municipios.csv"

	// Parse the file
	r := csv.NewReader(csvFile)
	r.Comma = ';'

	counter := 1
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		DB.Create(&Municipios{id: record[0], municipio: record[1]})

		fmt.Println(counter)
		counter++
	}
	// err := insertEmpresas("data/empresas/empresas.csv")
	// if err != nil {
	// 	fmt.Println("Erro ao inserir empresas.", err)
	// }
}
