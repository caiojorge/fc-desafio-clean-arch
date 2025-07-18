# Desafio

```
Olá devs!
Agora é a hora de botar a mão na massa. 
Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.
```
## Docker
- A aplicação deve ser executada dentro do docker-compose, inclusive o app.
- https://github.com/caiojorge/fc-desafio-clean-arch

## Subir o server
```
make up
```
## Verificar se o DB foi criado
```
http://localhost:8282/?server=mysql&username=root&db=orders
```
## 1. Teste do endpoint rest
- no diretorio API, testar usando o arquivo list_orders.http

## 2. Teste do Graphql
- grpc_orders.http

## 3. Teste GRPC
```bash
./api/test_grpc_list_orders.sh
```


