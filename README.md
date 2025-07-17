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

# Novo Endpoint para Listagem de Ordens

## 🎯 Endpoint Implementado

**GET /orders** - Lista todas as ordens cadastradas

## 🚀 Como testar

### 1. Preparar o banco de dados

```bash
# Subir os containers do MySQL e RabbitMQ
make up

# Executar o script SQL para criar tabela e dados de exemplo
docker exec -i mysql mysql -uroot -proot < init.sql
```

### 2. Executar a aplicação

```bash
# Compilar e executar
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```

### 3. Testar os endpoints

#### Listar todas as ordens:
```bash
curl -X GET http://localhost:8000/orders
```

#### Criar uma nova ordem:
```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "order-4",
    "price": 150.0,
    "tax": 15.0
  }'
```

## 📋 Resposta esperada do GET /orders

```json
[
  {
    "id": "order-1",
    "price": 100.0,
    "tax": 10.0,
    "final_price": 110.0
  },
  {
    "id": "order-2",
    "price": 200.0,
    "tax": 20.0,
    "final_price": 220.0
  },
  {
    "id": "order-3",
    "price": 50.0,
    "tax": 5.0,
    "final_price": 55.0
  }
]
```

## 🏗️ Arquitetura implementada

### Camadas envolvidas:

1. **Handler Web** (`WebListOrderHandler`) - Camada de apresentação HTTP
2. **Use Case** (`ListOrderUseCase`) - Lógica de negócio
3. **Repository** (`OrderRepository`) - Acesso aos dados
4. **Entity** (`Order`) - Entidade de domínio

### Fluxo da requisição:

```
HTTP Request → WebListOrderHandler → ListOrderUseCase → OrderRepository → Database
```

## 🧪 Testes

Todos os testes estão passando:

```bash
# Executar todos os testes
go test ./... -v

# Executar apenas testes do handler web
go test ./internal/infra/web/... -v

# Executar apenas testes do use case
go test ./internal/usecase/... -v
```


