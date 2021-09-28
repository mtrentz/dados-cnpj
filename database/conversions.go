package database

import (
	"database/sql"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

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
		log.Fatal("Erro ao parsear valor em ", fieldName, " ", err)
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
		// Como tem uma que outra ocorrencia errada no BD, aqui só printo o erro.
		log.Println("Erro ao parsear valor em ", fieldName, " ", err)
		return sql.NullTime{}
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
		log.Fatal("Erro ao parsear valor em ", fieldName, " ", err)
	}
	res = int64(math.Round(val))

	return sql.NullInt64{
		Int64: res,
		Valid: true,
	}
}
