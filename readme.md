<h1 align="center">
    <a href="https://pt-br.reactjs.org/">:moneybag: Simple Bank :moneybag:</a>
</h1>
<p align="center">🚀 Simple Bank é uma simples API REST que permite contas com saldo e transferencias entre si</p>

<div>
    <h2>🤖 Tecnologias</h2>
    <p><li>Golang</li> <li>SQLITE</li></p>
</div>

<div>
    <h2>📚 Bibliotecas</h2>
    <li><a href="github.com/dgrijalva/jwt-go">JWT-Go</a></li>
    <li><a href="github.com/google/uuid">Uuid</a></li>  
    <li><a href="github.com/gorilla/mux">Gorilla Mux</a></li>
    <li><a href="github.com/mattn/go-sqlite3">Go-Sqlite3</a></li>
    <li><a href="golang.org/x/crypto">Go Cryptography</a></li>
</div>

## Como rodar

#### Digite no terminal
```json
go run *.go
```
#### Utilize o PostMan e Dbeaver para facilitar


# :pushpin: Endpoints


## `/accounts`

### `localhost:8080/accounts - POST` Cria uma nova conta (Account). Exemplo:

```json
{
  "cpf": "075.139.133-03",
  "secret": "12345678",
  "name": "Gaab"
}
```
### `localhost:8080/accounts - GET` Mostra todas as contas (Accounts)

### `localhost:8080/accounts/{account_id}/balance - GET` Expõe o saldo (Balance) da conta (Account)

### Detalhes

- Somente Accounts não necessita de login, as demais como descobrir o balance e listar as accounts necessita de login
- As novas contas começam com 1000 no balance
- O saldo da conta sempre é inteiro
- De uma olhada [neste site](https://www.4devs.com.br/gerador_de_cpf) para gerar `cpf` aleatório
- O `secret` pode possuir de 6 a 30 caracteres

## `/login`

### `localhost:8080/login - POST` Cria um token JWT. Exemplo:
```json
// Request body
{
	"cpf": "075.139.133-03",
  "secret": "12345678"
}
// Response
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcGYiOiIwNzUuMTM5LjEzMy0wMyIsIm9yaWdpbmlkIjoiYzNhODc1MzAtZDcwYS00ZWExLWEyZTUtYTU4YTAyNTljOTU5IiwiZXhwIjoxNjMyNzk0MTI4fQ.W18x0GlKYX9VKb7FD22CN1_p7aI2jGTcTSwv-XGHPAM"
}
```

## `/transfer`

### `localhost:8080/transfer - POST` Transfere a outra conta. Exemplo:

```json
// Request body
{
  "destin": "d7a1bc2d-a702-45c2-8bd1-09671e203883",
  "amount": 240
}
```

### `localhost:8080/transfer - GET` Disponibiliza a lista de transferências realizadas

### Detalhes

- O TOKEN é necessário para solicitar a entrada
- A `account_origin_id` é acoplada ao TOKEN e `account_destination_id` é necessário para encontrar o destinatário do qual será direcionado a quantia 
- Para o usuário realizar transferências necessita de saldo em seu balance
