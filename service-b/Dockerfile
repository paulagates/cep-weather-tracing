FROM golang:1.23.2 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags='-w -s' -o /service-b cmd/main.go

FROM scratch

WORKDIR /

COPY --from=builder /service-b /service-b

EXPOSE 8081

CMD ["/service-b"]