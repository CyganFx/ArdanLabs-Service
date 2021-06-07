#FROM golang AS builder
#WORKDIR /app
#COPY go.mod ./
#RUN go mod download
#COPY . .
#
#RUN CGO_ENABLED=0 GOOS=linux go build -o myApp ./app/sales-api
#
#FROM alpine:latest
#WORKDIR /root
#RUN apk --no-cache add ca-certificates
#COPY --from=builder /app .
#
