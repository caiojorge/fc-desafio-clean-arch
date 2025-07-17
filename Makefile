up: 
	docker-compose up -d --build

down:
	docker-compose down

logs:
	docker-compose logs -f

logs-app:
	docker-compose logs -f app

restart:
	docker-compose restart app

build:
	docker-compose build

clean:
	docker-compose down -v
	docker system prune -f

status:
	docker-compose ps

run-wire: 
	wire ./cmd/ordersystem/

install-wire:
	go install github.com/google/wire/cmd/wire@latest

test:
	go test ./... -v

dev:
	go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

gen-proto:
	cd internal/infra/grpc/protofiles && protoc --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative order.proto

run-wire: 
	cd cmd/ordersystem && go generate .	
	cd cmd/ordersystem && wire
	wire cmd/ordersystem/wire.go
	cd cmd/ordersystem && go mod tidy && wire

test-grpc:
	grpcurl -plaintext -import-path internal/infra/grpc/protofiles -proto order.proto localhost:50051 pb.OrderService/ListOrders

run-py:
	ps aux | grep -E "(python|grpc_list_orders)" | grep -v grep
	lsof -i :50051

gen-grahql:
	go run github.com/99designs/gqlgen generate

test-graphql:
	curl -X POST http://localhost:8080/query \
  -H "Content-Type: application/json" \
  -d '{"query": "query { orders { id Price Tax FinalPrice } }"}'	