SHELL := /bin/bash

# Testing running system
# hey -m GET -c 100 -n 1000000 "http://localhost:3000/readiness"

# // To generate a private/public key PEM file.
# openssl genpkey -algorithm RSA -out private.pem -pkeyopt rsa_keygen_bits:2048
# openssl rsa -pubout -in private.pem -out public.pem

tidy:
	go mod tidy
	go mod vendor

run:
	go run ./app/sales-api/main.go

runa:
	go run ./app/sales-admin/main.go

monitor:
	expvarmon -ports=":4000" -vars="build,requests,goroutines,errors,mem:memstats.Alloc"