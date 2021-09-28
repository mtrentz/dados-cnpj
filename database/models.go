package database

import (
	"database/sql"
)

type Empresas struct {
	ID                         int64
	Cnpj                       string
	Razao_social               sql.NullString
	Id_natureza_juridica       sql.NullInt64
	Id_qualificacao            sql.NullInt64
	Capital_social             sql.NullInt64
	Id_porte                   sql.NullInt64
	Ente_federativo_resposavel sql.NullString
}

type Estabelecimentos struct {
	Cnpj                         string
	Cnpj_ordem                   string
	Cnpj_digito_verificador      string
	Id_matriz_filial             sql.NullInt64
	Nome_fantasia                sql.NullString
	Id_situacao_cadastral        sql.NullInt64
	Data_situacao_cadastral      sql.NullTime
	Id_motivo_situacao_cadastral sql.NullInt64
	Nome_cidade_exterior         sql.NullString
	Id_pais                      sql.NullInt64
	Data_inicio_atividade        sql.NullTime
	Id_cnae_principal            sql.NullString
	Lista_cnaes_secundarias      sql.NullString
	Tipo_logradouro              sql.NullString
	Logradouro                   sql.NullString
	Numero                       sql.NullString
	Complemento                  sql.NullString
	Bairro                       sql.NullString
	Cep                          sql.NullString
	Uf                           sql.NullString
	Id_municipio                 sql.NullInt64
	Ddd1                         sql.NullString
	Telefone1                    sql.NullString
	Ddd2                         sql.NullString
	Telefone2                    sql.NullString
	Ddd_fax                      sql.NullString
	Fax                          sql.NullString
	Email                        sql.NullString
	Situacao_especial            sql.NullString
	Data_situacao_especial       sql.NullTime
}

type Simples struct {
	Cnpj                     string
	Opcao_pelo_simples       sql.NullString
	Data_opcao_pelo_simples  sql.NullTime
	Data_exclusao_do_simples sql.NullTime
	Opcao_pelo_mei           sql.NullString
	Data_opcao_pelo_mei      sql.NullTime
	Data_entrada_do_mei      sql.NullTime
}

type Socios struct {
	Cnpj                                string
	Id_tipo_socio                       sql.NullInt64
	Nome_razao_social                   sql.NullString
	Cpf_cnpj                            sql.NullString
	Id_qualificacao                     sql.NullInt64
	Data_entrada                        sql.NullTime
	Id_pais                             sql.NullInt64
	Cpf_representante_legal             sql.NullString
	Nome_representante_legal            sql.NullString
	Id_qualificacao_representante_legal sql.NullInt64
	Id_faixa_etaria                     sql.NullInt64
}

type Paises struct {
	ID   int64
	Pais string
}

type QualificacoesDeSocios struct {
	ID                  string
	QualificacaoDeSocio string
}

type NaturezasJuridicas struct {
	ID               string
	NaturezaJuridica string
}

type Cnaes struct {
	Cnae          string
	CnaeDescricao string
}

type Motivos struct {
	ID                      string
	MotivoSituacaoCadastral string
}

type Municipios struct {
	ID        int64
	Municipio string
}
