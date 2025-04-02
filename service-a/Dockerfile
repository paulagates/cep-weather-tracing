FROM golang:1.23.2 as builder


WORKDIR /app

COPY ./ ./

RUN env CGO_ENABLED=0 && go mod tidy && go build -o service-a cmd/main.go

FROM scratch

WORKDIR /

COPY --from=builder /app/service-a ./

EXPOSE 8080

CMD ["./service-a"]