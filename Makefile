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