package organize

import (
	"testing"
)

func TestGetFilesAmount(t *testing.T) {
	fs := GetFiles("data")

	var tests = []struct {
		name           string
		pathList       []string
		expectedAmount int
	}{
		{"empresa", fs.Empresa, 10},
		{"estabelecimento", fs.Estabelecimento, 10},
		{"socio", fs.Socio, 10},
		{"simples", fs.Simples, 1},
		{"cnae", fs.Cnae, 1},
		{"motivo", fs.Motivo, 1},
		{"municipio", fs.Municipio, 1},
		{"natureza", fs.Natureza, 1},
		{"pais", fs.Pais, 1},
		{"qualificacao", fs.Qualificacao, 1},
	}

	for _, test := range tests {
		if foundAmount := len(test.pathList); foundAmount != test.expectedAmount {
			t.Errorf("TEST FAILED for %s. Expected %d; Got %d", test.name, test.expectedAmount, foundAmount)
		}
	}
}

// go test -v github.com/cuducos/minha-receita/organize -run TestGetFiles
