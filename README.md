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
- Fiz alguns ajustes e testei novamente em computadores diferentes (mac e no wsl)

- em caso de problemas com o mysql (porque mudei a versão)
```
# Parar todos os containers
docker-compose down

# Remover dados corrompidos do MySQL
sudo rm -rf .docker/mysql

# Limpar volumes do Docker
docker volume prune -f
```

## Subir o server
```
make up
```

## Desligar o server
```
make down
```

## Verificar se o DB foi criado
- acesso ao adminer
```
http://localhost:8282/?server=mysql&username=root&db=orders
```
## 1. Teste do endpoint rest
- no diretório api: list_orders.http

## 2. Teste do Graphql
- no diretório api: grafhql_orders.http

## 3. Teste GRPC
- o teste deve ser feito na raiz do projeto.
```bash
./api/test_grpc_list_orders.sh
```


