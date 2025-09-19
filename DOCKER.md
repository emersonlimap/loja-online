# 🐳 Docker Setup - Loja Online

Este guia explica como executar a aplicação Loja Online usando Docker.

## 📋 Pré-requisitos

- Docker
- Docker Compose

## 🚀 Como executar

### 1. Clone o repositório
```bash
git clone <url-do-repositorio>
cd loja-online
```

### 2. Execute com Docker Compose
```bash
# Build e start dos containers
docker-compose up --build

# Ou para rodar em background
docker-compose up --build -d
```

### 3. Acesse a aplicação
- **Aplicação**: http://localhost:8080
- **Login**: http://localhost:8080/login
- **Adminer (DB Manager)**: http://localhost:8081

### 4. Credenciais padrão
- **Email**: leozinsurfwear@gmail.com
- **Senha**: leozin@123

## 🛠️ Comandos úteis

### Ver logs
```bash
# Logs da aplicação
docker-compose logs app

# Logs do banco
docker-compose logs postgres

# Logs em tempo real
docker-compose logs -f
```

### Parar containers
```bash
docker-compose down
```

### Resetar tudo (remove volumes)
```bash
docker-compose down -v
docker-compose up --build
```

### Executar comandos no container
```bash
# Acessar o container da aplicação
docker-compose exec app sh

# Acessar o PostgreSQL
docker-compose exec postgres psql -U postgres -d loja_online
```

## 🗂️ Estrutura dos Containers

### 📱 App Container
- **Porta**: 8080
- **Imagem**: Construída a partir do Dockerfile
- **Volumes**: Logs em `/app/logs`

### 🗄️ PostgreSQL Container
- **Porta**: 5432
- **Versão**: PostgreSQL 15 Alpine
- **Database**: loja_online
- **Usuário**: postgres
- **Senha**: postgres

### 🔧 Adminer Container
- **Porta**: 8081
- **Função**: Interface web para gerenciar o banco

## 🔐 Configurações de Segurança

Para produção, altere:

1. **JWT Secret** no docker-compose.yml
2. **Senha do banco** no docker-compose.yml
3. **Use HTTPS** com proxy reverso (nginx/traefik)

## 🧪 Dados de Teste

O banco é inicializado com:
- ✅ Usuário admin padrão
- ✅ Produtos de exemplo da LeoZin Surfwear
- ✅ Cliente de exemplo
- ✅ Estoque inicial

## 🔍 Health Checks

Os containers incluem health checks:
- **App**: Verifica se a aplicação responde na porta 8080
- **DB**: Verifica se o PostgreSQL está aceitando conexões

## 📊 Monitoramento

Para monitorar os containers:
```bash
# Status dos containers
docker-compose ps

# Uso de recursos
docker stats

# Health status
docker-compose exec app wget --spider http://localhost:8080/
```