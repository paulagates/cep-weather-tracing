FROM golang:1.23.2 as builder

WORKDIR /app

COPY ./ ./

RUN env CGO_ENABLED=0 && go mod tidy && go build -o service-b cmd/main.go

FROM scratch

WORKDIR /

COPY --from=builder /app/service-b ./

EXPOSE 8081

CMD ["./service-b"]