# Variables
ENV_FILE = .env
ENV_SAMPLE_FILE = .env.example
DOCKER_COMPOSE = docker compose
DOCKER_COMPOSE_CMD = $(DOCKER_COMPOSE) up --build

all: copy-env up 

copy-env:
	@if [ ! -f $(ENV_FILE) ]; then \
		cp $(ENV_SAMPLE_FILE) $(ENV_FILE); \
		echo "$(ENV_FILE) created from $(ENV_SAMPLE_FILE)"; \
	else \
		echo "$(ENV_FILE) already exists"; \
	fi

# Start Docker containers
up:
	$(DOCKER_COMPOSE_CMD)

# Stop Docker containers
down:
	$(DOCKER_COMPOSE) down

# Rebuild Docker images and start containers
restart: down up

# Remove all stopped containers and dangling images
clean:
	$(DOCKER_COMPOSE) down --rmi all --volumes --remove-orphans

test:
	go test -v ./...

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans: 
	evans --host localhost --port 9090 -r repl

  .PHONY: all copy-env up down restart clean proto evans