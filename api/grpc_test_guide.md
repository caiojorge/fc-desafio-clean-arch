# Teste gRPC - Listagem de Ordens

## Pré-requisitos
- Servidor gRPC rodando na porta 50051
- Evans instalado (cliente gRPC)

## Instalação do Evans
```bash
# No Ubuntu/Debian
sudo apt install evans

# Ou via Go
go install github.com/ktr0731/evans@latest
```

## Testando a listagem de ordens

### 1. Conectar ao servidor gRPC
```bash
evans --host localhost --port 50051 --proto internal/infra/grpc/protofiles/order.proto
```

### 2. Listar ordens
```bash
# No prompt do Evans
call ListOrders
```

### 3. Criar uma ordem (para testar com dados)
```bash
# No prompt do Evans
call CreateOrder
# Então digite:
# id: "123"
# price: 100.0
# tax: 10.0
```

### 4. Listar ordens novamente
```bash
call ListOrders
```

## Usando grpcurl (alternativa)

### Instalar grpcurl
```bash
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```

### Testar listagem
```bash
grpcurl -plaintext -import-path internal/infra/grpc/protofiles -proto order.proto localhost:50051 pb.OrderService/ListOrders
```

### Criar ordem
```bash
grpcurl -plaintext -import-path internal/infra/grpc/protofiles -proto order.proto -d '{"id": "123", "price": 100.0, "tax": 10.0}' localhost:50051 pb.OrderService/CreateOrder
```
