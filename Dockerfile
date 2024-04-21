FROM golang:1.22.2-alpine3.19

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

COPY ./ ./

RUN go mod tidy

ENTRYPOINT CompileDaemon --build="go build -o build/go-api-test-task" -command="./build/go-api-test-task" -build-dir=/app