# Loja Online

Sistema de gerenciamento de loja online de roupas desenvolvido em Go com Gin e GORM.

## Funcionalidades

- ✅ Cadastro e gerenciamento de produtos
- ✅ Cadastro e gerenciamento de clientes  
- ✅ Sistema de vendas com controle de estoque
- ✅ Controle de inventário com movimentações
- ✅ Sistema de autenticação JWT
- ✅ Gerenciamento de usuários e permissões
- ✅ Relatórios de vendas
- ✅ Interface web básica

## Estrutura do Projeto

```
├── main.go                    # Ponto de entrada da aplicação
├── go.mod                     # Dependências Go
├── .env                       # Variáveis de ambiente
├── internal/
│   ├── api/
│   │   └── router.go         # Configuração das rotas
│   ├── config/
│   │   └── config.go         # Configurações da aplicação
│   ├── database/
│   │   └── database.go       # Conexão e migrations
│   ├── handlers/             # Handlers das rotas
│   │   ├── handler.go        # Handler base
│   │   ├── auth.go           # Autenticação
│   │   ├── products.go       # Produtos
│   │   ├── customers.go      # Clientes
│   │   ├── sales.go          # Vendas
│   │   ├── inventory.go      # Estoque
│   │   └── web.go            # Páginas web
│   ├── middleware/           # Middleware
│   │   ├── auth.go           # Autenticação JWT
│   │   ├── cors.go           # CORS
│   │   └── logging.go        # Logging
│   └── models/               # Modelos de dados
│       ├── customer.go       # Cliente
│       ├── inventory.go      # Estoque
│       ├── product.go        # Produto
│       ├── sale.go           # Venda
│       └── user.go           # Usuário
└── web/
    ├── static/
    │   ├── css/
    │   │   └── style.css     # Estilos
    │   └── js/
    │       └── app.js        # JavaScript
    └── templates/
        ├── dashboard.html    # Dashboard
        └── login.html        # Login
```

## Pré-requisitos

- Go 1.21+
- PostgreSQL
- Git

## Configuração

### 1. Clone o repositório
```bash
git clone <url-do-repositorio>
cd loja-online
```

### 2. Configure o banco de dados
```sql
CREATE DATABASE loja_online;
CREATE USER usuario WITH PASSWORD 'senha';
GRANT ALL PRIVILEGES ON DATABASE loja_online TO usuario;
```

### 3. Configure as variáveis de ambiente
```bash
cp .env.example .env
```

Edite o arquivo `.env` com suas configurações:
```env
PORT=8080
ENVIRONMENT=development
DATABASE_URL=postgres://usuario:senha@localhost:5432/loja_online?sslmode=disable
JWT_SECRET=sua_chave_secreta_jwt_aqui
```

### 4. Instale as dependências
```bash
go mod tidy
```

### 5. Execute a aplicação
```bash
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`

## Build para Produção

```bash
go build -o ./bin/loja-online .
./bin/loja-online
```

## API Endpoints

### Autenticação
- `POST /api/v1/auth/login` - Login
- `POST /api/v1/auth/register` - Registro

### Produtos (autenticação requerida)
- `GET /api/v1/products` - Listar produtos
- `POST /api/v1/products` - Criar produto
- `GET /api/v1/products/:id` - Obter produto
- `PUT /api/v1/products/:id` - Atualizar produto
- `DELETE /api/v1/products/:id` - Deletar produto

### Clientes (autenticação requerida)
- `GET /api/v1/customers` - Listar clientes
- `POST /api/v1/customers` - Criar cliente
- `GET /api/v1/customers/:id` - Obter cliente
- `PUT /api/v1/customers/:id` - Atualizar cliente
- `DELETE /api/v1/customers/:id` - Deletar cliente

### Vendas (autenticação requerida)
- `GET /api/v1/sales` - Listar vendas
- `POST /api/v1/sales` - Criar venda
- `GET /api/v1/sales/:id` - Obter venda
- `PUT /api/v1/sales/:id` - Atualizar venda

### Estoque (autenticação requerida)
- `GET /api/v1/inventory` - Listar inventário
- `POST /api/v1/inventory/adjust` - Ajustar estoque
- `GET /api/v1/inventory/movements/:product_id` - Movimentos de produto

### Usuários (autenticação requerida)
- `GET /api/v1/users` - Listar usuários
- `POST /api/v1/users` - Criar usuário
- `PUT /api/v1/users/:id` - Atualizar usuário
- `DELETE /api/v1/users/:id` - Deletar usuário

### Relatórios (autenticação requerida)
- `GET /api/v1/reports/sales` - Relatório de vendas

## Páginas Web
- `/` - Redirect para dashboard
- `/login` - Página de login
- `/dashboard` - Dashboard principal

## Tecnologias Utilizadas

- **Backend**: Go, Gin (framework web), GORM (ORM)
- **Banco de dados**: PostgreSQL
- **Autenticação**: JWT
- **Frontend**: HTML, CSS, JavaScript vanilla
- **Hash de senhas**: bcrypt

## Próximos Passos

- [ ] Implementar upload de imagens para produtos
- [ ] Adicionar filtros e busca avançada
- [ ] Implementar sistema de categorias
- [ ] Adicionar mais relatórios
- [ ] Melhorar interface web
- [ ] Adicionar testes unitários
- [ ] Implementar cache Redis
- [ ] Adicionar documentação Swagger

## Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.