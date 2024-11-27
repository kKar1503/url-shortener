FROM golang:1.23.1 as builder

WORKDIR /app

COPY go.mod ./

RUN go mod download;

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go;

FROM alpine:3.18 as final

ARG PORT=8080

WORKDIR /app
COPY --from=builder /app/main .

ENV PORT $PORT
EXPOSE $PORT

ENTRYPOINT ["./main"]
