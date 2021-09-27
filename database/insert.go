package database

func InsertAll() {

	DB.AutoMigrate(&Municipios{})

	// err := insertEmpresas("data/empresas/empresas.csv")
	// if err != nil {
	// 	fmt.Println("Erro ao inserir empresas.", err)
	// }
}
