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
)

func InsertAll() {
	err := insertEmpresas("data/empresas/empresas.csv")
	if err != nil {
		fmt.Println("Erro ao inserir empresas.", err)
	}
}

// TODO: Fazer testes agora q setei o autocommit como 0.... Talvez nao precise usar transactions!

func insertEmpresas(csvPath string) error {
	// Open file
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}

	// Parse the file
	r := csv.NewReader(csvFile)
	r.Comma = ';'

	// Begin the transaction
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal("EMPRESAS: Erro iniciando trasação, ", err)
	}

	qry := `
		INSERT INTO empresas
			(cnpj,
			razao_social,
			id_natureza_juridica,
			id_qualificacao,
			capital_social,
			id_porte,
			ente_federativo_resposavel)
		VALUES (?,?,?,?,?,?,?)`

	counter := 1
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// Unpack e converte valores, usa um struct do sql pra poder ser nulo
		cnpj := record[0]
		razao_social := newNullString(record[1])
		id_natureza_juridica := stringToNullInt(record[2], "id_natureza_juridica")
		id_qualificacao := stringToNullInt(record[3], "id_qualificacao")
		capital_social := floatStringToNullInt(record[4], "capital_social")
		id_porte := stringToNullInt(record[5], "id_porte")
		ente_federativo_resposavel := newNullString(record[6])

		_, err = tx.Exec(qry, cnpj, razao_social, id_natureza_juridica, id_qualificacao, capital_social, id_porte, ente_federativo_resposavel)
		if err != nil {
			log.Fatal("EMPRESAS: Erro executando query, ", err)
		}

		if counter%10000 == 0 {
			fmt.Printf("EMPRESAS: Inserido %d linhas\n", counter)
			// Commita as N linhas até agora
			tx.Commit()

			// Inicia uma nova transação
			tx, err = DB.Begin()
			if err != nil {
				log.Fatal("EMPRESAS: Erro iniciando NOVA trasação, ", err)
			}
		}
		counter++
	}
	// Comita as linhas que sobraram, se sobraram...
	err = tx.Commit()
	if err != nil {
		fmt.Println("EMPRESAS: Nada mais a comitar.")
	}

	return nil
}

// func insertEstabelecimentos(csvPath string) error {
// 	// Open file
// 	csvFile, err := os.Open(csvPath)
// 	if err != nil {
// 		return err
// 	}

// 	// Parse the file
// 	r := csv.NewReader(csvFile)
// 	r.Comma = ';'

// 	// Inicia transação
// 	tx, err := DB.Begin()
// 	if err != nil {
// 		log.Fatal("ESTABELECIMENTOS: Erro iniciando trasação, ", err)
// 	}

// 	qry := `
// 		INSERT INTO empresas
// 			(cnpj
// 			cnpj_ordem
// 			cnpj_digito_verificador
// 			id_matriz_filial
// 			nome_fantasia
// 			id_situacao_cadastral
// 			data_situacao_cadastral
// 			id_motivo_situacao_cadastral
// 			nome_cidade_exterior
// 			id_pais
// 			data_inicio_atividade
// 			id_cnae_principal
// 			lista_cnaes_secundarias
// 			tipo_logradouro
// 			logradouro
// 			numero
// 			complemento
// 			bairro
// 			cep
// 			uf
// 			id_municipio
// 			ddd1
// 			telefone1
// 			ddd2
// 			telefone2
// 			ddd_fax
// 			fax
// 			email
// 			situacao_especial
// 			data_situacao_especial)
// 		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

// 	counter := 1
// 	for {

// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		cnpj := record[0]
// 		cnpj_ordem := record[1]
// 		cnpj_digito_verificador := record[2]
// 		id_matriz_filial := stringToNullInt(record[3])
// 		nome_fantasia := newNullString(record[4])
// 		id_situacao_cadastral := stringToNullInt(record[5])
// 		data_situacao_cadastral :=
// 		id_motivo_situacao_cadastral := stringToNullInt(record[7])
// 		nome_cidade_exterior := newNullString(record[8])
// 		id_pais := stringToNullInt(record[9])
// 		data_inicio_atividade :=
// 		id_cnae_principal := newNullString(record[11])
// 		lista_cnaes_secundarias := newNullString(record[12])
// 		tipo_logradouro := newNullString(record[13])
// 		logradouro := newNullString(record[14])
// 		numero := newNullString(record[15])
// 		complemento := newNullString(record[16])
// 		bairro := newNullString(record[17])
// 		cep := newNullString(record[18])
// 		uf := newNullString(record[19])
// 		id_municipio := stringToNullInt(record[20])
// 		ddd1 := newNullString(record[21])
// 		telefone1 := newNullString(record[22])
// 		ddd2 := newNullString(record[23])
// 		telefone2 := newNullString(record[24])
// 		ddd_fax := newNullString(record[25])
// 		fax := newNullString(record[26])
// 		email := newNullString(record[27])
// 		situacao_especial := newNullString(record[28])
// 		data_situacao_especial :=

// 		_, err = tx.Exec(qry, cnpj, razao_social, id_natureza_juridica, id_qualificacao, capital_social, id_porte, ente_federativo_resposavel)
// 		if err != nil {
// 			log.Fatal("Erro executando query: ", err)
// 		}

// 		if counter%10000 == 0 {
// 			fmt.Printf("Inserido: %d linhas\n", counter)
// 			tx.Commit()

// 			tx, err = DB.Begin()
// 			if err != nil {
// 				log.Fatal("Erro iniciando NOVA trasação: ", err)
// 			}
// 		}
// 		counter++
// 	}
// 	err = tx.Commit()
// 	if err != nil {
// 		fmt.Println("Nada mais a comitar.")
// 	}

// 	return nil
// }

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
		log.Fatal("Erro ao parsear valor em: ", fieldName, err)
	}

	return sql.NullInt64{
		Int64: res,
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
		log.Fatal("Erro ao parsear valor em: ", fieldName, err)
	}
	res = int64(math.Round(val))

	return sql.NullInt64{
		Int64: res,
		Valid: true,
	}
}
