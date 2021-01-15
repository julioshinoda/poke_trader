.EXPORT_ALL_VARIABLES:

PORT=9001
FAIR_INDEX=20
DATABASE_URL=postgres://user:secret@localhost:5433/trade?sslmode=disable
run: 
	go run cmd/server.go


test:
	go test -v ./...