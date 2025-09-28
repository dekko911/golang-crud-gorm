run: 
	@CompileDaemon -command="./what-crud"

migrate-up:
	@go run migrations/main.go

test:
	@go test -v ./...