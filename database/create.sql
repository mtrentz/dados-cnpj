-- Para rodar só o MySQL no docker
/*
docker run --rm -d -v mysql:/var/lib/mysql \
  -v mysql_config:/etc/mysql -p 3366:3306 \
  --name dados-cnpj \
  -e MYSQL_ROOT_PASSWORD=dadoscnpj \
  mysql:latest \
  --character-set-server=latin2 --collation-server=latin2_general_ci 
*/

-- docker exec -it dados-cnpj bash
-- mysql -u root -p

-- Template da receita: https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/arquivos/NOVOLAYOUTDOSDADOSABERTOSDOCNPJ.pdf

DROP DATABASE IF EXISTS dados_cnpj;

CREATE DATABASE IF NOT EXISTS dados_cnpj CHARACTER SET latin2 COLLATE latin2_general_ci;

USE dados_cnpj;

CREATE TABLE IF NOT EXISTS empresas (
    cnpj VARCHAR(8) NOT NULL,
    razao_social TEXT NULL,
    id_natureza_juridica TEXT NULL,
    id_qualificacao INT NULL,
    capital_social BIGINT NULL, -- Vou converter pra int no Golang
    id_porte INT NULL, -- Vou criar tabela dos portes
    ente_federativo_resposavel TEXT NULL
);

CREATE TABLE IF NOT EXISTS estabelecimentos (
    cnpj VARCHAR(8) NOT NULL,
    cnpj_ordem VARCHAR(4) NOT NULL,
    cnpj_digito_verificador VARCHAR(2) NOT NULL,
    id_matriz_filial INT NULL, -- Vou criar tabela matriz/filial
    nome_fantasia TEXT NULL,
    id_situacao_cadastral INT NULL, -- Vou criar tabela sit cadastral
    data_situacao_cadastral DATE NULL, -- Preciso dar parse na data no GO. Formato AAAAMMDD
    id_motivo_situacao_cadastral TEXT NULL, -- Já vem com uma tabela com os motivos, vou deixar em string
    nome_cidade_exterior TEXT NULL,
    id_pais TEXT NULL, -- Vem com tabela pra linkar. Preciso passar as duas pra INT.
    data_inicio_atividade DATE NULL, -- Dar parse em date AAAAMMDD
    id_cnae_principal VARCHAR(7) NULL, -- Vou deixar esse em string por facilidade de procurar cnae na internet caso necessario.
    lista_cnaes_secundarias TEXT NULL, -- Esse aqui nao vou mudar nada, mas é uma string de cnaes separado por vírgula.
    tipo_logradouro TEXT NULL,
    logradouro TEXT NULL,
    numero TEXT NULL,
    complemento TEXT NULL,
    bairro TEXT NULL,
    cep TEXT NULL,
    uf TEXT NULL,
    id_municipio TEXT NULL, -- Vem com tabela pra linkar. Preciso passar pra INT
    ddd1 TEXT NULL,
    telefone1 TEXT NULL,
    ddd2 TEXT NULL,
    telefone2 TEXT NULL,
    ddd_fax TEXT NULL,
    fax TEXT NULL,
    email TEXT NULL,
    situacao_especial TEXT NULL,
    data_situacao_especial DATE NULL -- Não achei nenhum exemplo, vou assumir q é AAAAMMDD
);

CREATE TABLE IF NOT EXISTS simples (
    cnpj VARCHAR(8) NOT NULL,
    opcao_pelo_simples VARCHAR(1) NULL,
    data_opcao_pelo_simples DATE NULL, -- Dar parse na data.
    data_exclusao_do_simples DATE NULL, -- Dar parse na data.
    opcao_pelo_mei VARCHAR(1) NULL,
    data_opcao_pelo_mei DATE NULL, -- Dar parse na data.
    data_entrada_do_mei DATE NULL -- Dar parse na data.
);

CREATE TABLE IF NOT EXISTS socios (
    cnpj VARCHAR(8) NOT NULL,
	id_tipo_socio INT NULL, -- Vou criar tabela tipo socio
	nome_razao_social TEXT NULL,
	cpf_cnpj TEXT NULL,
	id_qualificacao TEXT NULL, -- Já tem tabela de qualificacao, passar pra INT
	data_entrada DATE NULL, -- Dar parse
    id_pais TEXT NULL,
	cpf_representante_legal TEXT NULL,
	nome_representante_legal TEXT NULL,
	id_qualificacao_representante_legal TEXT NULL, -- Linka pra mesmo qualificacao de socio que o id acima
    id_faixa_etaria INT NULL -- Vou criar tabela faixa etaria
);

CREATE TABLE IF NOT EXISTS paises (
    id TEXT NOT NULL,
    pais TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS qualificacoes_de_socios (
    id TEXT NOT NULL, -- Id é texto aqui pq se não o ID 00 da problema.
    qualificacao_de_socio TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS naturezas_juridicas (
    id TEXT NOT NULL, -- Id em text pq tem 0000 também
    natureza_juridica TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS cnaes (
    cnae VARCHAR(7) NOT NULL, -- Deixar em string para facilitar a busca dos cnaes online
    cnae_descricao TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS motivos (
    id TEXT NOT NULL,
    motivo_situacao_cadastral TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS municipios (
    id TEXT NOT NULL, -- preciso passar pra INT
    municipio TEXT NOT NULL
);

-- Criando algumas tabelas que vou inserir valores manualmente

CREATE TABLE IF NOT EXISTS portes (
    id INT NOT NULL, -- preciso passar pra INT
    porte TEXT NOT NULL
);

INSERT INTO portes (id, porte)
    VALUES (1, 'N/I'), (2, 'Micro Empresa'), (3, 'Pequeno Porte'), (5, 'Outros');


CREATE TABLE matriz_filiais (
    id INT PRIMARY KEY,
    matriz_filial TEXT NOT NULL
);

INSERT INTO matriz_filiais (id, matriz_filial)
    VALUES (1, 'Matriz'), (2, 'Filial');



CREATE TABLE situacoes_cadastrais (
    id INT PRIMARY KEY,
    situacao_cadastral TEXT NOT NULL
);

INSERT INTO situacoes_cadastrais (id, situacao_cadastral)
    VALUES (1, 'Nula'), (2, 'Ativa'), (3, 'Suspensa'), (4, 'Inapta'), (8, 'Baixada');


CREATE TABLE tipos_socios (
    id INT PRIMARY KEY,
    tipo_socio TEXT NOT NULL
);

INSERT INTO tipos_socios (id, tipo_socio)
    VALUES (1, 'Pessoa Juridica'), (2, 'Pessoa Fisica'), (3, 'Estrangeiro');


CREATE TABLE faixas_etarias (
    id INT PRIMARY KEY,
    faixa_etaria TEXT NOT NULL
);

INSERT INTO faixas_etarias (id, faixa_etaria)
    VALUES (1, '0 a 12 anos'),
           (2, '13 a 20 anos'),
           (3, '21 a 30 anos'),
           (4, '31 a 40 anos'),
           (5, '41 a 50 anos'),
           (6, '51 a 60 anos'),
           (7, '61 a 70 anos'),
           (8, '71 a 80 anos'),
           (9, 'Maiores de 80'),
           (0, 'Nao se aplica');
