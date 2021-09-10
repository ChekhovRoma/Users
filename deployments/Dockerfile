FROM golang:1.16-alpine

WORKDIR /app

COPY ./ ./

# install psql
RUN apk update
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

RUN go mod vendor

RUN go build -o ./build/users

EXPOSE 8080

ENTRYPOINT ["./build/users"]

CMD tail -f /dev/null