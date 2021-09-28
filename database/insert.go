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
	"time"
)

const batchSize = 1000

func InsertAll() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	insertMunicipios()
	// 	wg.Done()
	// }()
	// go func() {
	// 	insertSimples()
	// 	wg.Done()
	// }()
	// wg.Wait()

	insertMunicipios()
	insertEmpresas()
	insertSimples()
	insertEstabelecimentos()
	insertSocios()
	insertPaises()
	insertQualificacoes()
	insertNaturezas()
	insertCnaes()
	insertMotivos()
}

func insertEmpresas() {
	csvPath := "data/empresas/empresas.csv"

	r := getCsvReader(csvPath)

	var empresas []Empresas
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Empresas: Inserido %d linhas\n", counter)
					result := DB.Create(&empresas)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				empresas = append(empresas, Empresas{
					Cnpj:                       record[0],
					Razao_social:               newNullString(record[1]),
					Id_natureza_juridica:       stringToNullInt(record[2], "Empresas: id_natureza_juridica"),
					Id_qualificacao:            stringToNullInt(record[3], "Empresas: id_qualificacao"),
					Capital_social:             floatStringToNullInt(record[4], "Empresas: capital_social"),
					Id_porte:                   stringToNullInt(record[5], "Empresas: id_porte"),
					Ente_federativo_resposavel: newNullString(record[6]),
				})
				counter++
			}

			if counter%10000 == 0 {
				fmt.Printf("Empresas: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&empresas)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			empresas = nil
		}
	}()
}

func insertEstabelecimentos() {
	csvPath := "data/estabelecimentos/estabelecimentos.csv"

	r := getCsvReader(csvPath)

	var estabelecimentos []Estabelecimentos
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Estabelecimentos: Inserido %d linhas\n", counter)
					result := DB.Create(&estabelecimentos)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				estabelecimentos = append(estabelecimentos, Estabelecimentos{
					Cnpj:                         record[0],
					Cnpj_ordem:                   record[1],
					Cnpj_digito_verificador:      record[2],
					Id_matriz_filial:             stringToNullInt(record[3], "Estabelecimentos: id_matriz_filial"),
					Nome_fantasia:                newNullString(record[4]),
					Id_situacao_cadastral:        stringToNullInt(record[5], "Estabelecimentos: Id_situacao_cadastral"),
					Data_situacao_cadastral:      stringToNullTime(record[6], "Estebelecimentos: Data_situacao_cadastral"),
					Id_motivo_situacao_cadastral: stringToNullInt(record[7], "Estabelecimentos: Id_motivo_situacao_cadastral"),
					Nome_cidade_exterior:         newNullString(record[8]),
					Id_pais:                      stringToNullInt(record[9], "Estabelecimentos: Id_pais"),
					Data_inicio_atividade:        stringToNullTime(record[10], "Estabelecimentos: Data_inicio_atividade"),
					Id_cnae_principal:            newNullString(record[11]),
					Lista_cnaes_secundarias:      newNullString(record[12]),
					Tipo_logradouro:              newNullString(record[13]),
					Logradouro:                   newNullString(record[14]),
					Numero:                       newNullString(record[15]),
					Complemento:                  newNullString(record[16]),
					Bairro:                       newNullString(record[17]),
					Cep:                          newNullString(record[18]),
					Uf:                           newNullString(record[19]),
					Id_municipio:                 stringToNullInt(record[20], "Estabelecimentos: Id_municipio"),
					Ddd1:                         newNullString(record[21]),
					Telefone1:                    newNullString(record[22]),
					Ddd2:                         newNullString(record[23]),
					Telefone2:                    newNullString(record[24]),
					Ddd_fax:                      newNullString(record[25]),
					Fax:                          newNullString(record[26]),
					Email:                        newNullString(record[27]),
					Situacao_especial:            newNullString(record[28]),
					Data_situacao_especial:       stringToNullTime(record[29], "Estabelecimentos: Data_situacao_especial"),
				})
				counter++
			}

			if counter%10000 == 0 {
				fmt.Printf("Estabelecimentos: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&estabelecimentos)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			estabelecimentos = nil
		}
	}()
}

func insertSimples() {
	csvPath := "data/simples/simples.csv"

	r := getCsvReader(csvPath)

	var simples []Simples
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Simples: Inserido %d linhas\n", counter)
					result := DB.Create(&simples)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				simples = append(simples, Simples{
					Cnpj:                     record[0],
					Opcao_pelo_simples:       newNullString(record[1]),
					Data_opcao_pelo_simples:  stringToNullTime(record[2], "Simples: data_opcao_pelo_simples"),
					Data_exclusao_do_simples: stringToNullTime(record[3], "Simples: data_exclusao_do_simples"),
					Opcao_pelo_mei:           newNullString(record[4]),
					Data_opcao_pelo_mei:      stringToNullTime(record[5], "Simples: data_opcao_pelo_mei"),
					Data_entrada_do_mei:      stringToNullTime(record[6], "Simples: data_entrada_do_mei"),
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Simples: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&simples)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			simples = nil
		}
	}()
}

func insertSocios() {
	csvPath := "data/socios/socios.csv"

	r := getCsvReader(csvPath)

	var socios []Socios
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Socios: Inserido %d linhas\n", counter)
					result := DB.Create(&socios)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				socios = append(socios, Socios{
					Cnpj:                                record[0],
					Id_tipo_socio:                       stringToNullInt(record[1], "Socios: Id_tipo_socio"),
					Nome_razao_social:                   newNullString(record[2]),
					Cpf_cnpj:                            newNullString(record[3]),
					Id_qualificacao:                     stringToNullInt(record[4], "Socios: Id_qualificacao"),
					Data_entrada:                        stringToNullTime(record[5], "Socios: Data_entrada"),
					Id_pais:                             stringToNullInt(record[6], "Socios: Id_pais"),
					Cpf_representante_legal:             newNullString(record[7]),
					Nome_representante_legal:            newNullString(record[8]),
					Id_qualificacao_representante_legal: stringToNullInt(record[9], "Socios: Id_qualificacao_representante_legal"),
					Id_faixa_etaria:                     stringToNullInt(record[10], "Socios: Id_faixa_etaria"),
				})
				counter++
			}

			if counter%10000 == 0 {
				fmt.Printf("Socios: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&socios)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			socios = nil
		}
	}()
}

func insertPaises() {
	csvPath := "data/paises/paises.csv"

	r := getCsvReader(csvPath)

	var paises []Paises
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Paises: Inserido %d linhas\n", counter)
					result := DB.Create(&paises)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				paises = append(paises, Paises{
					ID:   stringToInt64(record[0], "Paises: id"),
					Pais: record[1],
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Paises: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&paises)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			paises = nil
		}
	}()
}

func insertQualificacoes() {
	csvPath := "data/qualificacoes_de_socios/qualificacoes_de_socios.csv"

	// DB.AutoMigrate(&QualificacoesDeSocios{})

	r := getCsvReader(csvPath)

	var qualificacoes []QualificacoesDeSocios
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Qualificacoes: Inserido %d linhas\n", counter)
					result := DB.Create(&qualificacoes)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				qualificacoes = append(qualificacoes, QualificacoesDeSocios{
					ID:                  record[0],
					QualificacaoDeSocio: record[1],
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Qualificacoes: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&qualificacoes)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			qualificacoes = nil
		}
	}()
}

func insertNaturezas() {
	csvPath := "data/naturezas_juridicas/naturezas_juridicas.csv"

	// DB.AutoMigrate(&QualificacoesDeSocios{})

	r := getCsvReader(csvPath)

	var naturezas []NaturezasJuridicas
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
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

				naturezas = append(naturezas, NaturezasJuridicas{
					ID:               record[0],
					NaturezaJuridica: record[1],
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Naturezas Juridicas: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&naturezas)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			naturezas = nil
		}
	}()
}

func insertCnaes() {
	csvPath := "data/cnaes/cnaes.csv"

	// DB.AutoMigrate(&QualificacoesDeSocios{})

	r := getCsvReader(csvPath)

	var cnaes []Cnaes
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Cnaes: Inserido %d linhas\n", counter)
					result := DB.Create(&cnaes)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				cnaes = append(cnaes, Cnaes{
					Cnae:          record[0],
					CnaeDescricao: record[1],
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Cnaes: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&cnaes)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			cnaes = nil
		}
	}()
}

func insertMotivos() {
	csvPath := "data/motivos/motivos.csv"

	// DB.AutoMigrate(&QualificacoesDeSocios{})

	r := getCsvReader(csvPath)

	var motivos []Motivos
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Motivos: Inserido %d linhas\n", counter)
					result := DB.Create(&motivos)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				motivos = append(motivos, Motivos{
					ID:                      record[0],
					MotivoSituacaoCadastral: record[1],
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Motivos: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&motivos)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			motivos = nil
		}
	}()
}

func insertMunicipios() {
	csvPath := "data/municipios/municipios.csv"

	r := getCsvReader(csvPath)

	var municipios []Municipios
	counter := 0
	func() {
		for {
			for i := 0; i < batchSize; i++ {
				record, err := r.Read()
				if err == io.EOF {
					fmt.Printf("Municipios: Inserido %d linhas\n", counter)
					result := DB.Create(&municipios)
					if result.Error != nil {
						log.Fatal(result.Error)
					}
					return
				}
				if err != nil {
					log.Fatal(err)
				}

				municipios = append(municipios, Municipios{
					ID:        stringToInt64(record[0], "Municipio: id"),
					Municipio: record[1],
				})
				counter++
			}
			if counter%10000 == 0 {
				fmt.Printf("Municipios: Inserido %d linhas\n", counter)
			}
			result := DB.Create(&municipios)
			if result.Error != nil {
				log.Fatal(result.Error)
			}
			municipios = nil
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

func stringToInt64(s string, fieldName string) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal("Erro ao parsear valor em: ", err)
	}
	return res
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
