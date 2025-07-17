#!/bin/bash

echo "=== Testando gRPC - Listagem de Ordens ==="
echo

# Verifica se o grpcurl está instalado
if ! command -v grpcurl &> /dev/null; then
    echo "grpcurl não está instalado. Instalando..."
    go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
fi

# Diretório do arquivo proto
PROTO_DIR="internal/infra/grpc/protofiles"
PROTO_FILE="order.proto"
SERVER="localhost:50051"

echo "1. Testando listagem de ordens (pode estar vazia)..."
grpcurl -plaintext -import-path $PROTO_DIR -proto $PROTO_FILE $SERVER pb.OrderService/ListOrders

echo
echo "2. Criando uma ordem de exemplo..."
grpcurl -plaintext -import-path $PROTO_DIR -proto $PROTO_FILE \
  -d '{"id": "order-123", "price": 100.0, "tax": 10.0}' \
  $SERVER pb.OrderService/CreateOrder

echo
echo "3. Listando ordens novamente..."
grpcurl -plaintext -import-path $PROTO_DIR -proto $PROTO_FILE $SERVER pb.OrderService/ListOrders

echo
echo "4. Criando outra ordem..."
grpcurl -plaintext -import-path $PROTO_DIR -proto $PROTO_FILE \
  -d '{"id": "order-456", "price": 200.0, "tax": 20.0}' \
  $SERVER pb.OrderService/CreateOrder

echo
echo "5. Listagem final..."
grpcurl -plaintext -import-path $PROTO_DIR -proto $PROTO_FILE $SERVER pb.OrderService/ListOrders

echo
echo "=== Teste concluído ==="
