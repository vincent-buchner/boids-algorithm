FROM golang:1.22.1-alpine3.19

WORKDIR /app

COPY go.mod go.sum main.go ./

RUN go mod download

COPY . .

EXPOSE 8080

CMD ["go", "run", "github.com/hajimehoshi/wasmserve@latest", "."]
