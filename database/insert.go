package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const batchSize = 1000

func InsertAll() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		insertEmpresas("data/empresas/empresas.csv")
		fmt.Println(">>> Empresas finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertEstabelecimentos("data/estabelecimentos/estabelecimentos.csv")
		fmt.Println(">>> Estabelecimentos finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertSimples("data/simples/simples.csv")
		fmt.Println(">>> Simples finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertSocios("data/socios/socios.csv")
		fmt.Println(">>> Socios finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertPaises("data/paises/paises.csv")
		fmt.Println(">>> Paises finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertQualificacoes("data/qualificacoes_de_socios/qualificacoes_de_socios.csv")
		fmt.Println(">>> Qualificacoes finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertCnaes("data/cnaes/cnaes.csv")
		fmt.Println(">>> Cnaes finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertMotivos("data/motivos/motivos.csv")
		fmt.Println(">>> Motivos finalizado.")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		insertMunicipios("data/municipios/municipios.csv")
		fmt.Println(">>> Municipios finalizado.")
		wg.Done()
	}()

	wg.Wait()

	// insertEmpresas("data/empresas/empresas.csv")
	// insertEstabelecimentos("data/estabelecimentos/estabelecimentos.csv")
	// insertSimples("data/simples/simples.csv")
	// insertSocios("data/socios/socios.csv")
	// insertPaises("data/paises/paises.csv")
	// insertQualificacoes("data/qualificacoes_de_socios/qualificacoes_de_socios.csv")
	// insertNaturezas("data/naturezas_juridicas/naturezas_juridicas.csv")
	// insertCnaes("data/cnaes/cnaes.csv")
	// insertMotivos("data/motivos/motivos.csv")
	// insertMunicipios("data/municipios/municipios.csv")
}

func insertEmpresas(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var empresas []Empresas
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {

				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Empresas: Inserido %d linhas\n", counter)
					result := DB.Create(&empresas)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				e := Empresas{}
				e.ReadRecord(record)
				empresas = append(empresas, e)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&empresas)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			empresas = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Empresas: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertEstabelecimentos(csvPath string) {
	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var estabelecimentos []Estabelecimentos
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {

				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Estabelecimentos: Inserido %d linhas\n", counter)
					result := DB.Create(&estabelecimentos)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				e := Estabelecimentos{}
				e.ReadRecord(record)
				estabelecimentos = append(estabelecimentos, e)
				counter++
			}

			// Insere o batch completo
			result := DB.Create(&estabelecimentos)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			estabelecimentos = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Estabelecimentos: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertSimples(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var simples []Simples
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {

				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Simples: Inserido %d linhas\n", counter)
					result := DB.Create(&simples)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				s := Simples{}
				s.ReadRecord(record)
				simples = append(simples, s)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&simples)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			simples = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Simples: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertSocios(csvPath string) {

	r := getCsvReader(csvPath)

	var socios []Socios
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {

				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Socios: Inserido %d linhas\n", counter)
					result := DB.Create(&socios)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				s := Socios{}
				s.ReadRecord(record)
				socios = append(socios, s)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&socios)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			socios = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Socios: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertPaises(csvPath string) {

	r := getCsvReader(csvPath)

	var paises []Paises
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {

				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Paises: Inserido %d linhas\n", counter)
					result := DB.Create(&paises)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				p := Paises{}
				p.ReadRecord(record)
				paises = append(paises, p)

				counter++
			}
			// Insere o batch completo
			result := DB.Create(&paises)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			paises = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Paises: Inserido %d linhas\n", counter)
			}

		}
	}()
}

func insertQualificacoes(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var qualificacoes []QualificacoesDeSocios
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Qualificacoes: Inserido %d linhas\n", counter)
					result := DB.Create(&qualificacoes)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				q := QualificacoesDeSocios{}
				q.ReadRecord(record)
				qualificacoes = append(qualificacoes, q)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&qualificacoes)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			qualificacoes = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Qualificacoes: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertNaturezas(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var naturezas []NaturezasJuridicas
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Naturezas Juridicas: Inserido %d linhas\n", counter)
					result := DB.Create(&naturezas)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				n := NaturezasJuridicas{}
				n.ReadRecord(record)
				naturezas = append(naturezas, n)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&naturezas)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			naturezas = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Naturezas Juridicas: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertCnaes(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var cnaes []Cnaes
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Cnaes: Inserido %d linhas\n", counter)
					result := DB.Create(&cnaes)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				c := Cnaes{}
				c.ReadRecord(record)
				cnaes = append(cnaes, c)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&cnaes)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			cnaes = nil

			if counter%10000 == 0 {
				fmt.Printf("Cnaes: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertMotivos(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var motivos []Motivos
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Motivos: Inserido %d linhas\n", counter)
					result := DB.Create(&motivos)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				m := Motivos{}
				m.ReadRecord(record)
				motivos = append(motivos, m)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&motivos)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			motivos = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Motivos: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func insertMunicipios(csvPath string) {

	// Pega o CSV Reader
	r := getCsvReader(csvPath)

	var municipios []Municipios
	counter := 0

	func() {
		// Loop infinito
		for {
			// Loop para ler os dados de UM BATCH
			for i := 0; i < batchSize; i++ {

				record, err := r.Read()
				if err == io.EOF {
					// Insere os dados que "sobram" do batch incompleto
					fmt.Printf("Municipios: Inserido %d linhas\n", counter)
					result := DB.Create(&municipios)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					// Sai totalmente da funcao
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				m := Municipios{}
				m.ReadRecord(record)
				municipios = append(municipios, m)

				counter++
			}

			// Insere o batch completo
			result := DB.Create(&municipios)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			// Limpa o slice do batch
			municipios = nil

			// Printa a cada 10mil inserções
			if counter%10000 == 0 {
				fmt.Printf("Municipios: Inserido %d linhas\n", counter)
			}
		}
	}()
}

func getCsvReader(csvPath string) *csv.Reader {
	// Open file
	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the file
	r := csv.NewReader(csvFile)
	r.Comma = ';'

	return r
}

func newNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func stringToNullInt(s string, fieldName string) sql.NullInt64 {
	if s == "" {
		return sql.NullInt64{}
	}

	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal("Erro ao parsear valor em: ", fieldName, " ", err)
	}

	return sql.NullInt64{
		Int64: res,
		Valid: true,
	}
}

func stringToNullTime(s string, fieldName string) sql.NullTime {
	if s == "" || s == "00000000" || s == "0" {
		return sql.NullTime{}
	}

	var res time.Time
	var err error

	// AAAAMMDD
	format := "20060102"

	res, err = time.Parse(format, s)
	if err != nil {
		log.Fatal("Erro ao parsear valor em: ", fieldName, " ", err)
	}
	return sql.NullTime{
		Time:  res,
		Valid: true,
	}

}

// Pega uma string tipo "10,00", passa pra int64 nullable.
func floatStringToNullInt(s string, fieldName string) sql.NullInt64 {
	if s == "" {
		return sql.NullInt64{}
	}

	var res int64

	s = strings.Replace(s, ",", ".", -1)
	val, err := strconv.ParseFloat(s, 32)
	if err != nil {
		log.Fatal("Erro ao parsear valor em: ", fieldName, " ", err)
	}
	res = int64(math.Round(val))

	return sql.NullInt64{
		Int64: res,
		Valid: true,
	}
}
