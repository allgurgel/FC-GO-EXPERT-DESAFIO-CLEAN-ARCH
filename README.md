
<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://plataforma.fullcycle.com.br/static/media/logo.6d87ce09.svg" alt="Project logo"></a>
</p>

<h1 align="center">Full Cycle Go Expert</h1>

---

<p align="center"> Desafio Curso GO Expert 
    <br> 
</p>

## 🧐 About <a name = "about"></a>

Clean Architecture
Requisitos:

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar a listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)

- Service ListOrders com GRPC

- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

## ⛏️ Built Using <a name = "built_using"></a>

- [Go]() - Golang
- [gqlgen](https://gqlgen.com/) - gqlgen

   
## 🏁 Getting Started <a name = "getting_started"></a>

Instruções para rodar o projeto.

```
docker-compose up -d

make migrate

go mod tidy
cd cmd/ordersystem

go run main.go wire_gen.go
```


---
### Prerequisites

Software que você precisa instalar.

- [Go](https://go.dev/dl/) - 1.19
- [golang-migrate](https://github.com/golang-migrate/migrate) - CLI
- [Docker]()

---
### Others


Schema GraphQL
```
mutation createOrder {
  createOrder(input:{
    id:"ccc",
    Price:10.5
    Tax: 1
  }){
    id,
    Price,
    Tax,
    FinalPrice
  }
}

query listOrders {
  listOrders(input:{
    page:1,
    limit: 20,
    sort: "asc"
  }) {
    id,
    Price,
    Tax,
    FinalPrice
  }
}
```
Comando para gerar o graphql
```
go run github.com/99designs/gqlgen generate

```

Comando para gerar código com protofile

```
protoc --go_out=.   --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```




