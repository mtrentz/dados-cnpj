package cmd

import (
	"github.com/mtrentz/dados-cnpj/download"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Baixa os dados do site da receita federal.",
	Long: `Procura todos os arquivos zips contidos no site da receita federal
e baixa-os para a pasta data.

Os servidores da receita são devagares e instáveis. Assim, além do download demorar
bastante ele pode não completar.

O tamanho de todos os zips pesam em torno de 4 GB.

É recomendado entrar diretamente no link abaixo da receita, baixar todos os zips e coloca-los
na pasta data.

https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/dados-publicos-cnpj`,
	Run: func(cmd *cobra.Command, args []string) {
		download.DownloadAll("data")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
