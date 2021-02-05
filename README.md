**Digital Account API**
-----------------------

API de transferência entre contas internas de um banco digital.

### A API utiliza-se das seguintes tecnologias / pacotes / frameworks:

* Go v1.15
* Postgres
* Postgres Driver (package jackc/pgx)
* Swagger (package swag)
* JWT (JSON Web Token para autenticação)
* Docker
* Dockertest (para testes de repository)
* Travis (Teste e CD/CD)


### Dependências:

* [Docker](https://docs.docker.com/engine/install/)
* [Docker-Compose](https://docs.docker.com/compose/install/)


## Uso da API

Iniciar a API:

``` sh
make start
```

Desligar a API:

``` sh
make stop
```

Testar a API:

``` sh
make test
```

Build da API:

``` sh
make build
```


### API Request

| Metódo | Rota                           | Descrição                                        | Autenticação |
|--------|--------------------------------|--------------------------------------------------|--------------|
| POST   | /accounts                      | cria uma conta                                   | Não          |
| GET    | /accounts                      | obtém a lista de contas                          | Não          |
| GET    | /accounts/{account_id}/balance | obtém o saldo da conta                           | Não          |
| POST   | /login                         | autentica o usuário                              | Não          |
| GET    | /transfers                     | retorna as transferências da conta no banco      | Sim          |
| POST   | /transfers                     | faz transferencia de uma conta para outra        | Sim          |
| GET    | /health-check                  | health check                                     | Não          |


### Documentação completa da API (Redoc):

https://digital-accounts.herokuapp.com/redoc

ou local (api inicializada com os valores padrões):

http://localhost:9090/redoc


### Documentação interativa (Swagger):

https://digital-accounts.herokuapp.com

ou local (api inicializada com os valores padrões):

http://localhost:9090


### Configurações da API

A API faz uso de variáveis de ambiente.
Por padrão a API utiliza as variáveis de ambiente sitadas abaixo. Modifique apenas as que necessita alterar:

```
# Web Config
HOST="localhost"
PORT="9090"

# Database Config
DATABASE_HOST="localhost"
DATABASE_PORT="5439"
DATABASE_USER="postgres"
DATABASE_PASSWORD="postgres"
DATABASE_NAME="digital_accounts"
DATABASE_SSLMODE="disable"
DATABASE_POOL_MIN_SIZE="2"
DATABASE_POOL_MAX_SIZE="10"

# Authentication Config
JWT_SECRET_KEY="jwt-digital-accounts-secret-key"
JWT_MAX_TOKEN_LIVE_TIME="50m"

# API Config
PROFILE="dev"
MIGRATION_PATH="app/persistence/database/postgres/migrations"

```

### Desenvolvedor
- Bruno de Castro Oliveira - [brunodecastro](https://github.com/brunodecastro)

