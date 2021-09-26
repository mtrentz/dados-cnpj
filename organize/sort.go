package organize

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type FileNameLists struct {
	// Lista dos nomes dos arquivos de cada tipo
	Empresa         []string
	Estabelecimento []string
	Socio           []string
	Simples         []string
	Cnae            []string
	Motivo          []string
	Municipio       []string
	Natureza        []string
	Pais            []string
	Qualificacao    []string
}

func GetFiles(dataDir string) FileNameLists {

	organizedFiles := FileNameLists{}

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		// Quero só arquivos e não diretorios
		if f.IsDir() {
			continue
		}

		fName := f.Name()
		// Procura pela string que identifica o tipo de cada arquivo ZIP
		if strings.Contains(strings.ToUpper(fName), "EMPRECSV.ZIP") {
			organizedFiles.Empresa = append(organizedFiles.Empresa, fName)
		} else if strings.Contains(strings.ToUpper(fName), "ESTABELE.ZIP") {
			organizedFiles.Estabelecimento = append(organizedFiles.Estabelecimento, fName)
		} else if strings.Contains(strings.ToUpper(fName), "SOCIOCSV.ZIP") {
			organizedFiles.Socio = append(organizedFiles.Socio, fName)
		} else if strings.Contains(strings.ToUpper(fName), "SIMPLES.CSV") {
			organizedFiles.Simples = append(organizedFiles.Simples, fName)
		} else if strings.Contains(strings.ToUpper(fName), "CNAECSV.ZIP") {
			organizedFiles.Cnae = append(organizedFiles.Cnae, fName)
		} else if strings.Contains(strings.ToUpper(fName), "MOTICSV.ZIP") {
			organizedFiles.Motivo = append(organizedFiles.Motivo, fName)
		} else if strings.Contains(strings.ToUpper(fName), "MUNICCSV.ZIP") {
			organizedFiles.Municipio = append(organizedFiles.Municipio, fName)
		} else if strings.Contains(strings.ToUpper(fName), "NATJUCSV.ZIP") {
			organizedFiles.Natureza = append(organizedFiles.Natureza, fName)
		} else if strings.Contains(strings.ToUpper(fName), "PAISCSV.ZIP") {
			organizedFiles.Pais = append(organizedFiles.Pais, fName)
		} else if strings.Contains(strings.ToUpper(fName), "QUALSCSV.ZIP") {
			organizedFiles.Qualificacao = append(organizedFiles.Qualificacao, fName)
		} else {
			fmt.Printf("Arquivo %s não foi reconhecido como sendo parte de uma categoria\n", fName)
		}
	}

	// fmt.Printf("%+v\n", organizedFiles)

	return organizedFiles

}
