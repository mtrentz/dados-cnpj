package transform

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Aqui quero fazer uma funcao que pega as pastas lá do data/categoria,
// vai juntando cada csv em um só e dai já deleta eles

func GetFilesByCategory(dataDir string) {
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	// Dando loop pelos diretorios na pasta do path
	for _, f := range files {
		// Aqui quero só diretorios e não arquivos
		if !f.IsDir() {
			continue
		}
		dirName := f.Name()
		fmt.Printf("Dir Name: %s, dir Files:\n", dirName)

		dirFiles, err := ioutil.ReadDir(path.Join(dataDir, dirName))
		if err != nil {
			log.Fatal(err)
		}

		for _, df := range dirFiles {
			fileName := df.Name()
			fmt.Printf(">>> %s\n", fileName)
		}

	}
}

func ConcatCsvs() {
	dir := "data/Socio/K3241.K03200Y0.D10911.SOCIOCSV"

	// File Read
	inputFile, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()
	r := csv.NewReader(inputFile)
	r.Comma = ';'

	// File Write
	outputFile, err := os.Create("data/Socio/socios.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer outputFile.Close()
	w := csv.NewWriter(outputFile)
	w.Comma = ';'

	counter := 0
	for {
		if counter >= 10 {
			break
		}

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		counter++
		fmt.Println(record)
		w.Write(record)
	}

	w.Flush()
}
