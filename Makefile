tidy:
	go mod tidy
	go mod vendor

run:
	go run ./app/sales-api/main.go