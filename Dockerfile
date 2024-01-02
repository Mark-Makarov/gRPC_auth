FROM golang:1.21.5

WORKDIR /app

RUN apt-get update && apt-get install -y gcc sqlite3

COPY go.mod .
COPY go.sum .

COPY . .

WORKDIR /app
RUN go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations

WORKDIR /app/cmd/sso

RUN go build -o server

EXPOSE 6868:6868

CMD ["./server"]