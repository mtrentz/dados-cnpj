package cmd

import (
	"github.com/mtrentz/dados-cnpj/organize"
	"github.com/mtrentz/dados-cnpj/transform"
	"github.com/spf13/cobra"
)

// transformCmd represents the transform command
var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "Unzip todos os arquivos baixados da receita federal. Concatena todos os csvs em um só arquivo.",
	Long: `Tendo todos os arquivos zips da receita federal baixados na pasta data, unzipa todos eles
para uma pasta com o nome de sua categoria (empresa, cnae, etc...). Como a receita providencia vários csvs,
aqui também todos estes arquivos são concatenados em um único csv para cada categoria.

Para garantir o bom funcionamento dessa etapa remova todas as pastas dentro da pasta data e coloque todos os
arquivos zip com o mesmo nome dos arquivos baixados do site da receita federal.`,
	Run: func(cmd *cobra.Command, args []string) {
		OrganizedFiles := organize.GetFiles("data")
		transform.UnzipAll(OrganizedFiles, "data")
		transform.ConcatAll("data")
	},
}

func init() {
	rootCmd.AddCommand(transformCmd)
}
