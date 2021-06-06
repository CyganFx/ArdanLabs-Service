SHELL := /bin/bash

# Testing running system
# hey -m GET -c 100 -n 1000000 "http://localhost:3000/readiness"

tidy:
	go mod tidy
	go mod vendor

run:
	go run ./app/sales-api/main.go

monitor:
	expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,mem:memstats.Alloc"