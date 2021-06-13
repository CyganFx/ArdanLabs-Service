FROM golang AS builder

RUN mkdir /app
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o myApp ./app/sales-api/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app .

CMD ["./myApp"]

LABEL org.opencontainers.image.title="sales-api" \
      org.opencontainers.image.authors="Duman Ishanov <duman070601@gmail.com>" \
      org.opencontainers.image.source="https://github.com/CyganFx/ArdanLabs-Service" \
      org.opencontainers.image.vendor="CyganFx"