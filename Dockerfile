FROM golang:1.24.4-alpine

WORKDIR /app
COPY . .

RUN go build -o api ./src/cmd/api

EXPOSE 8080
CMD ["./api"]