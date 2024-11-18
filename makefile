proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans: 
	evans --host localhost --port 9090 -r repl

postgres: 
	docker run --name pg-db-con -p 5439:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg-db-con createdb --username=root --owner=root grpc_project_db

dropdb:
	docker exec -it pg-db-con dropdb grpc_project_db

server: 
	go run main.go

  .PHONY: proto evans postgres createdb dropdb server