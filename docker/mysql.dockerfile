FROM mysql:latest

ADD ./database/create.sql /docker-entrypoint-initdb.d