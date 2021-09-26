package organize

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func GetFiles(dataDir string) map[string][]string {
	OrganizedFiles := make(map[string][]string)

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
			OrganizedFiles["empresas"] = append(OrganizedFiles["empresas"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "ESTABELE.ZIP") {
			OrganizedFiles["estabelecimentos"] = append(OrganizedFiles["estabelecimentos"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "SOCIOCSV.ZIP") {
			OrganizedFiles["socios"] = append(OrganizedFiles["socios"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "SIMPLES.CSV") {
			OrganizedFiles["simples"] = append(OrganizedFiles["simples"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "CNAECSV.ZIP") {
			OrganizedFiles["cnaes"] = append(OrganizedFiles["cnaes"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "MOTICSV.ZIP") {
			OrganizedFiles["motivos"] = append(OrganizedFiles["motivos"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "MUNICCSV.ZIP") {
			OrganizedFiles["municipios"] = append(OrganizedFiles["municipios"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "NATJUCSV.ZIP") {
			OrganizedFiles["naturezas_juridicas"] = append(OrganizedFiles["naturezas_juridicas"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "PAISCSV.ZIP") {
			OrganizedFiles["paises"] = append(OrganizedFiles["paises"], fName)
		} else if strings.Contains(strings.ToUpper(fName), "QUALSCSV.ZIP") {
			OrganizedFiles["qualificacoes_de_socios"] = append(OrganizedFiles["qualificacoes_de_socios"], fName)
		} else {
			fmt.Printf("Arquivo %s não foi reconhecido como sendo parte de uma categoria\n", fName)
		}
	}

	return OrganizedFiles
}
