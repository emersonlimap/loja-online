# Multi-stage build para otimizar o tamanho da imagem
FROM golang:1.24.7-alpine AS builder

# Definir variáveis de ambiente para Go
ENV GOPROXY=direct
ENV GOSUMDB=off

# Definir diretório de trabalho
WORKDIR /app

# Instalar dependências necessárias para o Go baixar módulos
RUN apk add --no-cache git

# Copiar tudo primeiro
COPY . .

# Baixar dependências e fazer build
RUN go mod download || echo "Falha no download, tentando build direto..." && \
    CGO_ENABLED=0 GOOS=linux go build -o main . || \
    (echo "Build falhou, listando arquivos:" && ls -la && exit 1)

# Stage final - usar scratch para imagem mínima
FROM scratch

# Copiar binário da aplicação do stage builder
COPY --from=builder /app/main /main

# Copiar arquivos estáticos e templates
COPY --from=builder /app/web /web

# Expor porta
EXPOSE 8080

# Comando para executar a aplicação
CMD ["/main"]