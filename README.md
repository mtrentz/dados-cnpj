# Dados-Cnpj

Programa para baixar, limpar e importar os dados abertos de CNPJ para um banco de MySQL.

O banco está conforme ao [novo padrão da receita federal](https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/dados-publicos-cnpj).

## Requisitos

O mais fácil é rodar utilizando o docker, no qual será criado um container contendo o banco em MySQL, acessável no port 3307.

Porém é possível rodar com Go e um banco MySQL para conectar.

## Rodando com Docker
```
docker volume create cnpj_db
docker-compose build

docker-compose run --rm dados-cnpj download   // Downloads podem ser instáveis, ler a seção de Dados no README
docker-compose run --rm dados-cnpj transform
docker-compose run --rm dados-cnpj insert
```

## Rodando sem Docker

Pra rodar sem docker é necessário configurar algum banco de dados, preferencialmente MySQL, alterando a variável dsn em database/database.go.

```
{USER}:{PASSWD}@tcp({HOST}:{PORT})/{DB_NAME}?charset=latin2&collation=latin2_general_ci&autocommit=false&parseTime=true
```

Além disso é necessário preparar o banco a partir do dados/create.sql.

Por fim:
```
go get

go run main.go download
go run main.go transform
go run main.go insert
```

## Outros databases

A conexão do Go é feita utilizando [GORM](https://gorm.io/index.html), assim, é possível apenas alterar o arquivo database.go e importar
o driver correto para conectar com outros bancos como PostgreSQL ou SQLite. Preparando o banco conforme o create.sql é provavel que a importação funcione sem maiores problemas.

## Dados

O download dos [dados](https://www.gov.br/receitafederal/pt-br/assuntos/orientacao-tributaria/cadastros/consultas/dados-publicos-cnpj) providenciados pela
receita federal pode ser bastante instável. As velocidades são bem baixas e muitas vezes não concluêm. Assim, **recomendo baixar diretamente pelo site** ao invés
de utilizar o comando download, garantindo que todos os arquivos zips foram baixados por completos.

Após isso basta colocar os arquivos ZIP dentro da pasta data e rodar os comandos transform e insert.

## Sobre

Esse projeto foi bastante baseado na implementação do [minha-receita](https://github.com/cuducos/minha-receita) do Cuducos, porém funcionando com a nova versão dos dados e com o MySQL.



