**Digital Account API**
-----------------------

API de transferência entre contas internas de um banco digital.

### A API utiliza-se das seguintes tecnologias / pacotes / frameworks:

* Go v1.15 [golang](https://golang.org/)
* Postgres [postgres](https://www.postgresql.org/)
* Postgres Driver [jackc/pgx](https://github.com/jackc/pgx)
* Migração de Banco de Dados [golang-migrate/migrate](github.com/golang-migrate/migrate)
* Configuração de variáveis de ambiente [kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig)
* HTTP router [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
* Logs [uber-go/zap](https://github.com/uber-go/zap)
* Swagger [swaggo/swag](https://github.com/swaggo/swag)
* JWT (JSON Web Token para autenticação) [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
* Docker [Docker](https://docs.docker.com/engine/install/)
* Dockertest (para testes de repository) [ory/dockertest](https://github.com/ory/dockertest/v3)
* Travis (Teste e CI/CD) [travis](https://travis-ci.org/)
* Heroku (Deploy da aplicação) [heroku](http://www.heroku.com)


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

Logs da API:

``` sh
make logs
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
SWAGGER_HOST="localhost:9090"

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
EXECUTE_MIGRATION=true
MIGRATION_PATH="app/persistence/database/postgres/migrations"

```

### Desenvolvedor
- Bruno de Castro Oliveira - [brunodecastro](https://github.com/brunodecastro)

