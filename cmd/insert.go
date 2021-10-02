package cmd

import (
	"github.com/mtrentz/dados-cnpj/database"
	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insere os arquivos csvs concatenados ao database",
	Long: `Insere os arquivos csvs de cada categoria já concatenados e localizados dentro
da pasta data/categoria ao banco de dados.

É necessário manter a estrutura dos diretórios como é criado pelo comando transform. Ou seja,
dentro da pata data deve haver outras pastas com os nome das categorias dos arquivos: cnaes, empresas, etc...

Dentro dessas pastas deve haver um único csv que será inserido ao banco.

As configurações do banco de dados ficam em database/database.go. Neste arquivo é possível
mudar o local do banco de destino caso seja necessário.`,
	Run: func(cmd *cobra.Command, args []string) {
		database.Connect()
		database.InsertAll()
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
}
