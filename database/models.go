package database

import (
	"database/sql"
)

type Empresas struct {
	Id_cnpj                    int64
	Cnpj                       string
	Razao_social               sql.NullString
	Id_natureza_juridica       sql.NullInt64
	Id_qualificacao            sql.NullInt64
	Capital_social             sql.NullInt64
	Id_porte                   sql.NullInt64
	Ente_federativo_resposavel sql.NullString
}

// Lê uma linha do CSV e faz as conversões dos dados
func (e *Empresas) ReadRecord(record []string) {
	e.Id_cnpj = stringToInt64(record[0], "Empresas: id_cnpj")
	e.Cnpj = record[0]
	e.Razao_social = newNullString(record[1])
	e.Id_natureza_juridica = stringToNullInt(record[2], "Empresas: id_natureza_juridica")
	e.Id_qualificacao = stringToNullInt(record[3], "Empresas: id_qualificacao")
	e.Capital_social = floatStringToNullInt(record[4], "Empresas: capital_social")
	e.Id_porte = stringToNullInt(record[5], "Empresas: id_porte")
	e.Ente_federativo_resposavel = newNullString(record[6])
}

type Estabelecimentos struct {
	Id_cnpj                      int64
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

// Lê uma linha do CSV e faz as conversões dos dados
func (e *Estabelecimentos) ReadRecord(record []string) {
	e.Id_cnpj = stringToInt64(record[0], "Estabelecimentos: id_cnpj")
	e.Cnpj = record[0]
	e.Cnpj_ordem = record[1]
	e.Cnpj_digito_verificador = record[2]
	e.Id_matriz_filial = stringToNullInt(record[3], "Estabelecimentos: id_matriz_filial")
	e.Nome_fantasia = newNullString(record[4])
	e.Id_situacao_cadastral = stringToNullInt(record[5], "Estabelecimentos: Id_situacao_cadastral")
	e.Data_situacao_cadastral = stringToNullTime(record[6], "Estebelecimentos: Data_situacao_cadastral")
	e.Id_motivo_situacao_cadastral = stringToNullInt(record[7], "Estabelecimentos: Id_motivo_situacao_cadastral")
	e.Nome_cidade_exterior = newNullString(record[8])
	e.Id_pais = stringToNullInt(record[9], "Estabelecimentos: Id_pais")
	e.Data_inicio_atividade = stringToNullTime(record[10], "Estabelecimentos: Data_inicio_atividade")
	e.Id_cnae_principal = newNullString(record[11])
	e.Lista_cnaes_secundarias = newNullString(record[12])
	e.Tipo_logradouro = newNullString(record[13])
	e.Logradouro = newNullString(record[14])
	e.Numero = newNullString(record[15])
	e.Complemento = newNullString(record[16])
	e.Bairro = newNullString(record[17])
	e.Cep = newNullString(record[18])
	e.Uf = newNullString(record[19])
	e.Id_municipio = stringToNullInt(record[20], "Estabelecimentos: Id_municipio")
	e.Ddd1 = newNullString(record[21])
	e.Telefone1 = newNullString(record[22])
	e.Ddd2 = newNullString(record[23])
	e.Telefone2 = newNullString(record[24])
	e.Ddd_fax = newNullString(record[25])
	e.Fax = newNullString(record[26])
	e.Email = newNullString(record[27])
	e.Situacao_especial = newNullString(record[28])
	e.Data_situacao_especial = stringToNullTime(record[29], "Estabelecimentos: Data_situacao_especial")
}

type Simples struct {
	Id_cnpj                  int64
	Cnpj                     string
	Opcao_pelo_simples       sql.NullString
	Data_opcao_pelo_simples  sql.NullTime
	Data_exclusao_do_simples sql.NullTime
	Opcao_pelo_mei           sql.NullString
	Data_opcao_pelo_mei      sql.NullTime
	Data_entrada_do_mei      sql.NullTime
}

// Lê uma linha do CSV e faz as conversões dos dados
func (s *Simples) ReadRecord(record []string) {
	s.Id_cnpj = stringToInt64(record[0], "Simples: id_cnpj")
	s.Cnpj = record[0]
	s.Opcao_pelo_simples = newNullString(record[1])
	s.Data_opcao_pelo_simples = stringToNullTime(record[2], "Simples: data_opcao_pelo_simples")
	s.Data_exclusao_do_simples = stringToNullTime(record[3], "Simples: data_exclusao_do_simples")
	s.Opcao_pelo_mei = newNullString(record[4])
	s.Data_opcao_pelo_mei = stringToNullTime(record[5], "Simples: data_opcao_pelo_mei")
	s.Data_entrada_do_mei = stringToNullTime(record[6], "Simples: data_entrada_do_mei")
}

type Socios struct {
	Id_cnpj                             int64
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

// Lê uma linha do CSV e faz as conversões dos dados
func (s *Socios) ReadRecord(record []string) {
	s.Id_cnpj = stringToInt64(record[0], "Socios: id_cnpj")
	s.Cnpj = record[0]
	s.Id_tipo_socio = stringToNullInt(record[1], "Socios: Id_tipo_socio")
	s.Nome_razao_social = newNullString(record[2])
	s.Cpf_cnpj = newNullString(record[3])
	s.Id_qualificacao = stringToNullInt(record[4], "Socios: Id_qualificacao")
	s.Data_entrada = stringToNullTime(record[5], "Socios: Data_entrada")
	s.Id_pais = stringToNullInt(record[6], "Socios: Id_pais")
	s.Cpf_representante_legal = newNullString(record[7])
	s.Nome_representante_legal = newNullString(record[8])
	s.Id_qualificacao_representante_legal = stringToNullInt(record[9], "Socios: Id_qualificacao_representante_legal")
	s.Id_faixa_etaria = stringToNullInt(record[10], "Socios: Id_faixa_etaria")
}

type Paises struct {
	ID   string
	Pais string
}

// Lê uma linha do CSV e faz as conversões dos dados
func (p *Paises) ReadRecord(record []string) {
	p.ID = record[0]
	p.Pais = record[1]
}

type QualificacoesDeSocios struct {
	ID                  string
	QualificacaoDeSocio string
}

// Lê uma linha do CSV e faz as conversões dos dados
func (q *QualificacoesDeSocios) ReadRecord(record []string) {
	q.ID = record[0]
	q.QualificacaoDeSocio = record[1]
}

type NaturezasJuridicas struct {
	ID               string
	NaturezaJuridica string
}

// Lê uma linha do CSV e faz as conversões dos dados
func (n *NaturezasJuridicas) ReadRecord(record []string) {
	n.ID = record[0]
	n.NaturezaJuridica = record[1]
}

type Cnaes struct {
	Cnae          string
	CnaeDescricao string
}

// Lê uma linha do CSV e faz as conversões dos dados
func (c *Cnaes) ReadRecord(record []string) {
	c.Cnae = record[0]
	c.CnaeDescricao = record[1]
}

type Motivos struct {
	ID                      string
	MotivoSituacaoCadastral string
}

// Lê uma linha do CSV e faz as conversões dos dados
func (m *Motivos) ReadRecord(record []string) {
	m.ID = record[0]
	m.MotivoSituacaoCadastral = record[1]
}

type Municipios struct {
	ID        string
	Municipio string
}

// Lê uma linha do CSV e faz as conversões dos dados
func (m *Municipios) ReadRecord(record []string) {
	m.ID = record[0]
	m.Municipio = record[1]
}
