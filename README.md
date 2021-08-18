# Tecnologias

* Go 1.16.5

* Docker 🐳

* MySQL 5.7 🐬

* Clean Architecture

  

# Bibliotecas

* [Migrate](https://github.com/golang-migrate/migrate)

* [Swaggo](https://github.com/swaggo/swag)

* [JWT-GO](https://github.com/dgrijalva/jwt-go)

* [MUX](https://github.com/gorilla/mux)

* [Satori UUID](https://github.com/satori/go.uuid)


# ChallengeGo

O ChallengeGo é uma API para um simples sistema bancário, com os seguintes objetivos:

* Criar uma conta;
* Autenticar-se;
* Realizar depósitos;
* Realizar transferências;
* Realizar saques;

A API acessa um banco de dados em MySQL, onde são mantidas duas tabelas:
* account
* transfer

# Como utilizar

O banco é versionado com migrate. Ao fazer o deploy da API, as tabelas serão criadas. Para alterar as tabelas, versionando o banco, podem ser executados os comandos:


```bash

migrate -path gateway/db/mysql/migration -database "mysql://root:password@tcp(localhost:3306)/sys?multiStatements=true" -verbose up

```

<h2>Rotas da aplicação</h2>

A API está documentada com swagger que se encontra na url http://localhost:8080/docs/swagger/

Exemplo de criação de conta:

http://localhost:8080/accounts<br>

Body json:

```javascript
{
    "nome": "Fran",
    "cpf":"98711222334",
    "secret": "123"  
}
```
Para efetuar login e gerar o token, que permitirá o acesso às rotas autenticadas (/transfer, /deposit, /withdraw), é preciso passar as credenciais, conforme o exemplo:

http://localhost:8080/auth<br>

```javascript
{
    "cpf":"98711222334",
    "secret": "123"  
}
```


A API acompanha um dockerfile e um docker-compose, que já faz o deploy de um banco mysql na versão 5.7 e da API na porta 8080. Para realizar tais ações, basta executar o comando:

```bash

docker-compose up

```

  

<h2> Devs </h2>

Existe no projeto um arquivo Make para facilitar o build e a execução da aplicação. Para utilizá-lo, basta executar os comandos:
  

```bash

make build

```

  
  

```bash

make run

```

  

```bash

make test

```


<h2> Environment variables </h2>

Executando a aplicação sem o docker, é necessário configurar as variáveis de ambiente em um arquivo .env na raíz do projeto. O nome a descrição de cada uma delas está a seguir na tabela.  

| Variável  |  Descrição  |
| ------------------- | ------------------- |
|  DB_NAME |  Nome do banco de dados |
|  DB_USER |  Usuário do banco |
|  DB_PASS |  Senha do banco   |
|  DB_HOST |  Endereço do servidor (ex: localhost)
|  DB_PORT |  Porta do banco (3306 - mysql)
