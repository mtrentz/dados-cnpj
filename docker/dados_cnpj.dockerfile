FROM golang:1.17.0 AS builder
WORKDIR /app
COPY . .
RUN apt-get update
RUN apt-get-upgrade -y
RUN CGO_ENABLED=0 GOOS=linux go mod tidy && go get && go build -o dados-cnpj main.go

FROM debian:buster-slim AS production
WORKDIR /app
COPY --from=builder /app/dados-cnpj /usr/bin/dados-cnpj
ENTRYPOINT ["/usr/bin/dados-cnpj"]