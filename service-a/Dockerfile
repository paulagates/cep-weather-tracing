FROM golang:1.23.2 as builder


WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -ldflags='-w -s' -o /service-a cmd/main.go

FROM scratch

COPY --from=builder /service-a /service-a

EXPOSE 8080

CMD ["/service-a"]