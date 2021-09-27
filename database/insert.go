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

func insertEmpresas(csvPath string) error {
	// Open file
	csvFile, err := os.Open(csvPath)
	if err != nil {
		return err
	}

	// Parse the file
	r := csv.NewReader(csvFile)
	r.Comma = ';'

	// Prepara o query de inserção
	og_stmt, err := DB.Prepare(`
			INSERT INTO empresas
				(cnpj,
				razao_social,
				id_natureza_juridica,
				id_qualificacao,
				capital_social,
				id_porte,
				ente_federativo_resposavel)
			VALUES (?,?,?,?,?,?,?)`)
	if err != nil {
		return err
	}
	defer og_stmt.Close()

	// Começa a primeira transação
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal("Erro ao começar transação", err)
	}
	insert_stmt := tx.Stmt(og_stmt)

	// Read lines
	fmt.Println("Inserindo dados das empresas...")
	counter := 1
	for {
		// A cada N inserções printo uma mensagem e comito a transação de inserção
		if counter%10000 == 0 {
			fmt.Printf("Empresas: Inserido %d linhas\n", counter)
			// Finaliza a transação
			tx.Commit()

			// Inicia uma nova transação
			tx, err := DB.Begin()
			if err != nil {
				log.Fatal("Erro ao começar transação", err)
			}
			insert_stmt = tx.Stmt(og_stmt)
		}
		counter++

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// TODO: Precisa dar um check ainda pra colocar em utf8 acho, algo assim...
		// Error 1366: Incorrect string value: '\xC3O' for column
		// Unpack e converte valores, usa um struct do sql pra poder ser nulo
		cnpj := record[0]
		razao_social := newNullString(record[1])
		id_natureza_juridica := stringToNullInt(record[2], "id_natureza_juridica")
		id_qualificacao := stringToNullInt(record[3], "id_qualificacao")
		capital_social := floatStringToNullInt(record[4], "capital_social")
		id_porte := stringToNullInt(record[5], "id_porte")
		ente_federativo_resposavel := newNullString(record[6])

		// Insere no database
		_, err = insert_stmt.Exec(cnpj, razao_social, id_natureza_juridica, id_qualificacao, capital_social, id_porte, ente_federativo_resposavel)
		if err != nil {
			return err
		}
	}
	// Commita o resto
	tx.Commit()
	return nil
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
