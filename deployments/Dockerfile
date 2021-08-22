FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./build/users

EXPOSE 8080

ENTRYPOINT ./build/users