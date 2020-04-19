# Make file

build:
	go build -o safe-cash-service main.go

run:
	go run main.go

stop-services:
	docker-compose -f ./setup/postgreSQL.yml down
	docker-compose -f ./setup/server.yml down
	docker volume rm docker-compose_database-data

local-db:
	docker-compose -f ./setup/postgreSQL.yml up -d
	@while ! docker exec safe-cash-service_postgres pg_isready -h localhost -p 5432 > /dev/null; do \
		sleep 1; \
	done
	# docker cp ./setup/docker-compose/data.sql cinema-admin_postgres:/
	# docker exec -u postgres safe-cash-service_postgres psql safe-cash-service user -f /data.sql

server:
	docker-compose -f ./setup/server.yml up -d

setup-package:
	mkdir logs
	go get github.com/Masterminds/glide
	glide install

test:
	go test -v ./...

# docker-image:
# 	make build 